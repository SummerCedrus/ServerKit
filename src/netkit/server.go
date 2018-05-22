package netkit

import (
	"fmt"
	"net"
)
var mgr *ConnectMgr
func NewServer(addr string) (*ConnectMgr, error) {
	mgr = &ConnectMgr{
		MsgChan: make(chan *Message, MaxMsgNum),
		ConnectChan: make(chan *Receiver, MaxSessionNum),
		CloseChan: make(chan bool),
		ReceiverPool:make(map[int32]*Receiver, MaxSessionNum),
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)

	if nil != err {
		return mgr, err
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)

	if nil != err {
		return mgr, err
	}

	go func() {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept failed: ", err)
				continue
			}
			NewReciever(conn)
		}
	}()

	return mgr, nil
}

func CloseServer() {
	close(mgr.MsgChan)
}
