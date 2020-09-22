package main

import (
	"encoding/binary"
	"errors"
	"net"
)

// GetIPNet returns the interfaces Network address information.
func GetIPNet(iface *net.Interface) (*net.IPNet, error) {
	var addr *net.IPNet
	// Get IP addresses from interface
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}
	// Iterate through addresses
	for _, a := range addrs {
		// Assert address of type IPNet and store
		if ipnet, ok := a.(*net.IPNet); ok {
			// Check type net.IP present and store
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				// Determine network mask
				mask := ipnet.Mask[len(ipnet.Mask)-4:]
				// Construct IPNet
				addr = &net.IPNet{
					IP:   ip4,
					Mask: mask,
				}
			}
		}
	}
	// Sanity checks
	if addr == nil {
		return nil, errors.New("no IPv4 address found")
	} else if addr.IP[0] == 127 {
		return nil, errors.New("skipping localhost")
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		return nil, errors.New("mask means network is too large")
	}
	return addr, nil
}

// HostIPGen returns a slice containing all possible IP addresses.
// within the network address's subnet
func HostIPGen(n *net.IPNet) ([]net.IP, error) {
	if n == nil || n.IP == nil {
		return nil, errors.New("invalid IPNet provided")
	}
	// Represent network Ip and Mask as 32bit
	ip := binary.BigEndian.Uint32([]byte(n.IP))
	// fmt.Println("ip bytes:", ip)
	// fmt.Printf("ip unint32: %b \n", ip)
	// fmt.Printf("ip type: %T\n", ip)
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	// Recovery netid by masking ip
	netip := (ip & mask)
	// Identify first hostid ip
	hostip := netip + 1
	// Incremement mask to find all possible hostid masks
	var hostips []net.IP
	for mask < 0xffffffff {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], hostip)
		hostips = append(hostips, net.IP(buf[:]))
		mask++
		hostip++
	}
	return hostips, nil
}
