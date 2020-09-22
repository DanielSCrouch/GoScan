package main

import (
	"fmt"
	"net"
	"testing"
)

func TestDatagram(t *testing.T) {
	iface := net.Interface{
		Index:        6,
		MTU:          1500,
		Name:         "en0",
		HardwareAddr: net.HardwareAddr{0x88, 0xe9, 0xfe, 0x86, 0x50, 0x98},
		Flags:        21,
	}
	fmt.Println(iface)
	dstip := net.IPv4(byte(10), byte(0), byte(20), byte(202)) //10.0.2.202
	srcip := net.IPv4(byte(10), byte(0), byte(20), byte(56))  //10.0.2.5
	datagram := NewArpRequest(dstip, srcip, &iface)
	datagramResult := []byte{255, 255, 255, 255, 255, 255, 136, 233, 254, 134, 80, 152, 8, 6, 0, 1, 8, 0, 6, 16, 0, 1, 136, 233, 254, 134, 80, 152, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 10, 0, 20, 56, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 10, 0, 20, 202}

	for i := range datagram {
		if datagram[i] != datagramResult[i] {
			t.Errorf("Mismatch")
		}
	}
}
