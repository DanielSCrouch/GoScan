package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creat a new type of 'host'
// which is a slice of strings

type hosts []string

func (h hosts) print() {
	for i, host := range h {
		fmt.Println(i, host)
	}
}

func (h hosts) getHost(pos int) string {
	return h[pos]
}

func (h hosts) toString() string {
	s := strings.Join(h, ",")
	return s
}

func (h hosts) saveToFile(filename string) error {
	e := ioutil.WriteFile(filename, []byte(h.toString()), 0666)
	return e
}

func getHostsFromFile(filename string) hosts {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return hosts(s)
}

func findHosts() hosts {
	// main function
	var hosta string = "a"
	hostb := "b"

	hosts := hosts{"d", "e"}

	hosts = append(hosts, hosta, hostb)

	return hosts
}

func (h hosts) randomSelect() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range h {
		newPosition := rand.Intn(len(h) - 1)
		newPosition = r.Intn(len(h) - 1)
		h[i], h[newPosition] = h[newPosition], h[i]
	}
}
