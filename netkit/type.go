package netkit

import (
	"net"
	"github.com/golang/protobuf/proto"
)

const (
	MaxSessionNum = 20
	MaxMsgNum = 20
)
var MaxSessionID = int32(0)
type ConnectMgr struct {
	MsgChan chan *Message
	CloseChan chan bool
	ConnectChan chan *Receiver
	ReceiverPool map[int32]*Receiver
}

type Receiver struct {
	ID		int32
	Conn    *net.TCPConn
	IP      net.IP
	MsgChan chan *Message
	//CloseChan chan bool
}

type Sender struct {
	Queue chan *Message
	Conn    *net.TCPConn
	CloseChan chan bool
}

type Message struct {
	Cmd uint32
	Msg proto.Message
}
