package main

import "fmt"

// Creat a new type of 'host'
// which is a slice of strings

type host struct {
	name string
	mac  string
	ipv4 string
}

func (hPointer *host) updateName(name string) {
	(*hPointer).name = name
}

func (h host) print() {
	fmt.Printf("%+v\n", h)
}
