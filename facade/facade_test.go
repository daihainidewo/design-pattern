// 外观模式（Facade Pattern）隐藏系统的复杂性，并向客户端提供了一个客户端可以访问系统的接口。
// 这种类型的设计模式属于结构型模式，它向现有的系统添加一个接口，来隐藏系统的复杂性。

package facade

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

type ConnectorMgr struct {
	tc TCPConnector
	uc UDPConnector
}

func (cm ConnectorMgr) ConnectTCP() {
	cm.tc.Connect()
}

func (cm ConnectorMgr) ConnectUDP() {
	cm.uc.Connect()
}

func TestFacade(t *testing.T) {
	cm := ConnectorMgr{
		tc: TCPConnector{},
		uc: UDPConnector{},
	}
	cm.ConnectTCP()
	cm.ConnectUDP()
}