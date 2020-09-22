package main

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// NewArpRequest returns an ARP Request datagram
func NewArpRequest(dstIP net.IP, srcIP net.IP, iface *net.Interface) []byte {
	// Sender protocol address
	ip := srcIP
	if srcIP == nil {
		ipnet, _ := GetIPNet(iface)
		ip = ipnet.IP
	}
	// Set up buffer and options for serialization.
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	// Ethernet layers
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	//  ARP packet layers
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,    // Hardware Type, RFC 826
		Protocol:          layers.EthernetTypeIPv4,    // Protocol Type, Typically IPv4 EtherType
		HwAddressSize:     6,                          // Hardware address Length of sender and receiver
		ProtAddressSize:   4,                          // IPv4 Protocol address length
		Operation:         layers.ARPRequest,          // ARP Request operation
		SourceHwAddress:   []byte(iface.HardwareAddr), // Sender hardware address
		SourceProtAddress: []byte(ip),                 // Sender protocol address
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},   // Destination protocol address
		DstProtAddress:    []byte(dstIP),              // Destination protocol address
	}
	// Serialise packet (marshalling)
	err := gopacket.SerializeLayers(buf, opts, &eth, &arp)
	if err != nil {
		panic(err)
	}
	// Return datagram
	datagram := buf.Bytes()
	return datagram
}

// RESOURCES
// https://en.wikipedia.org/wiki/Address_Resolution_Protocol
// https://github.com/j-keck/arping/blob/master/arp_datagram.go
// https://github.com/mdlayher/arp/blob/master/packet.go
