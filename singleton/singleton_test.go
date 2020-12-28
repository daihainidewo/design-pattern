package singleton

import (
	"fmt"
	"testing"
)

type TCPConnect struct {
}

var tc *TCPConnect

func GetTCPConnect() *TCPConnect {
	if tc == nil {
		return &TCPConnect{}
	}
	return tc
}

func TestSingleton(t *testing.T) {
	fmt.Printf("%p\n", tc)
	tc1 := GetTCPConnect()
	fmt.Printf("%p\n", tc1)
	tc2 := GetTCPConnect()
	fmt.Printf("%p\n", tc2)
}
