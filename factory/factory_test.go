package factory

import (
	"fmt"
	"testing"
)

type Connector interface {
	Connect()
}

type TCPConnector struct {
}

func (c TCPConnector) Connect() {
	fmt.Println("tcp connect")
}

type UDPConnector struct {
}

func (c UDPConnector) Connect() {
	fmt.Println("udp connect")
}

type ConnectFactory struct {
}

func (cf ConnectFactory) GetConnect(t string) Connector {
	switch t {
	case "TCP":
		return TCPConnector{}
	case "UDP":
		return UDPConnector{}
	}
	return nil
}

func TestFactory(t *testing.T) {
	cf := ConnectFactory{}
	tc := cf.GetConnect("TCP")
	uc := cf.GetConnect("UDP")

	tc.Connect()
	uc.Connect()
}
