package main

import "fmt"

// Host type
type Host string

// HostList type
type HostList []Host

// HostSet type
type HostSet map[Host]interface{}

// Add item to HostSet
func (s HostSet) Add(n Host) {
	s[n] = struct{}{}
}

// Remove item to HostSet
func (s HostSet) Remove(n Host) {
	delete(s, n)
}

// Contains item to HostSet
func (s HostSet) Contains(n Host) bool {
	_, found := s[n]
	return found
}

func main() {
	s := make(HostSet)
	s.Add("google.org")
	s.Add("google.com")
	s.Add("gopheracademy.org")
	s.Remove("google.com")

	hostnames := HostList{
		"google.org",
		"google.com",
		"gopheracademy.org",
	}

	for _, n := range hostnames {
		fmt.Printf("%s? %v\n", n, s.Contains(n))
	}
}
