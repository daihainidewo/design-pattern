// 装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。
// 这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

package decorator

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
	c Connector
}

func (cm ConnectorMgr) Connect() {
	cm.c.Connect()
}

func (cm ConnectorMgr) Close() {
	fmt.Println("connect close")
}

func TestDecorator(t *testing.T) {
	tc := ConnectorMgr{
		c: TCPConnector{},
	}
	uc := ConnectorMgr{
		c: UDPConnector{},
	}
	tc.Connect()
	tc.Close()
	uc.Connect()
	uc.Close()
}
