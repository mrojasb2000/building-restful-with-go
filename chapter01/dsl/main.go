package main

import "fmt"

import "strings"

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

// Select items on value equal on f(host)
func (l HostList) Select(f func(Host) bool) HostList {
	result := make(HostList, 0, len(l))
	for _, h := range l {
		if f(h) {
			result = append(result, h)
		}
	}
	return result
}

// IsDotOrg find item end with ".org"
func IsDotOrg(h Host) bool {
	return strings.HasSuffix(string(h), ".org")
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

	fmt.Printf("HostList: %v\n", hostnames.Select(IsDotOrg))
}
