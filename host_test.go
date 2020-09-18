package main

import "testing"

func TestHosts(t *testing.T) {
	h := findHosts()

	if len(h) != 4 {
		t.Errorf("Expected 4 hosts but got %d", len(h))
	}
}
