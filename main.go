package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Handle option parsing (Interface name, default en0 for mac)
	var interfaceName string
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "interfaceName")
		interfaceName = "en0"
	} else {
		interfaceName = os.Args[1]
	}

	// Get the hosts' interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	// Scan interface
	for _, iface := range ifaces {
		if iface.Name == interfaceName {
			hostips, err := arpScan(&iface)
			if err != nil {
				panic(err)
			}
			for _, hostip := range hostips {
				fmt.Printf("Host IP: %s \t Type: %T\n", hostip.String(), hostip)
			}
		}
	}
}

// arpScan identies live hosts by sending ARP requests to all hosts
// in the subnet before identifying host ARP responses.
func arpScan(iface *net.Interface) (hostips []net.IP, err error) {
	fmt.Println("Scanning on interface:", iface.Name)

	// Open a pcap handle for receiving live packets from interface
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	defer handle.Close() // Schedule closing handle

	// Start a goroutine to read in arp packet data
	fmt.Println("Polling for packets...")
	stop := make(chan bool) // channel for stop signal
	go arpRead(handle, iface, &hostips, stop)
	defer close(stop)

	// Write arp request packets to subnet addresses
	fmt.Println("Writing ARP Requests.")
	ip, _ := GetIPNet(iface)
	subnetips, _ := HostIPGen(ip)
	for i := 1; i < 5; i++ {
		for _, ip := range subnetips {
			arpWrite(handle, iface, ip)
		}
		time.Sleep(2 * time.Second)
	}
	// Wait for responses
	time.Sleep(2 * time.Second)
	// Stop polling for responses
	stop <- true
	// Return IP addresses of detected hosts
	return hostips, nil
}

// Read ARP responses from interface
func arpRead(handle *pcap.Handle, iface *net.Interface, hostips *[]net.IP, stop chan bool) {
	// Setup packet ethernet layer decoder
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	// Continuously poll for arp packets
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			fmt.Println("Stopped polling for packets.")
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)

			if arp.Operation != layers.ARPReply { // || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress)
				// The ARP Packet sent
				continue
			} else {
				hostip := net.IP(arp.SourceProtAddress)
				addUniqueIP(hostips, hostip)
				continue
			}
		}
	}
}

// Writes an ARP request to a host address via the interface
func arpWrite(handle *pcap.Handle, iface *net.Interface, dstIP net.IP) error {
	datagram := NewArpRequest(dstIP, nil, iface)
	if err := handle.WritePacketData(datagram); err != nil {
		return err
	}
	return nil
}

// Add ip to slice if unique address
func addUniqueIP(ips *[]net.IP, ip net.IP) {
	unique := true
	for _, v := range *ips {
		if v.Equal(ip) {
			unique = false
		}
	}
	if unique {
		*ips = append(*ips, ip)
	}
}

// RESOURCES

// https://github.com/google/gopacket/blob/master/examples/arpscan/arpscan.go
// https://godoc.org/github.com/google/gopacket/pcap
// https://github.com/hellojukay/arp-scanner/blob/master/main.go

// fmt.Println("hostip: ", hostip)
// fmt.Printf("%T\n", hostip)
