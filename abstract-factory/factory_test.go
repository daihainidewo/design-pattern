package abstract_factory

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
	AbstractFactory
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

type Closer interface {
	Close()
}

type TCPCloser struct {
}

func (c TCPCloser) Close() {
	fmt.Println("tcp close")
}

type UDPCloser struct {
}

func (c UDPCloser) Close() {
	fmt.Println("udp Close")
}

type CloseFactory struct {
	AbstractFactory
}

func (cf CloseFactory) GetCloser(t string) Closer {
	switch t {
	case "TCP":
		return TCPCloser{}
	case "UDP":
		return UDPCloser{}
	}
	return nil
}

type AbstractFactory interface {
	GetConnect(t string) Connector
	GetCloser(t string) Closer
}

type Factory struct {
	connectF ConnectFactory
	closeF   CloseFactory
}

func (f Factory) GetFactory(t string) AbstractFactory {
	switch t {
	case "CONNECT":
		return f.connectF
	case "CLOSE":
		return f.closeF
	}
	return nil
}

func TestFactory(t *testing.T) {
	f := Factory{}
	connectF := f.GetFactory("CONNECT")
	tc := connectF.GetConnect("TCP")
	uc := connectF.GetConnect("UDP")

	closeF := f.GetFactory("CLOSE")
	tclose := closeF.GetCloser("TCP")
	uclose := closeF.GetCloser("UDP")

	tc.Connect()
	uc.Connect()

	tclose.Close()
	uclose.Close()
}
