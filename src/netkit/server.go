package netkit

import (
	"fmt"
	//"net"
	"net"
	//"io/ioutil"
)
func NewServer(addr string) (*ClientMsg, error){
	msg := ClientMsg{
		MsgChan: make(chan [] byte, 5),
		//CloseChan: make(chan bool),
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)

	if nil != err{
		return &msg, err
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)

	if nil != err{
		return &msg, err
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept failed: ", err)
				continue
			}
			buff := make([]byte, 20)
			_, err = conn.Read(buff)

			if err != nil {
				fmt.Println("Read failed: ", err)
				continue
			}
			fmt.Println(buff)
			msg.MsgChan <- buff
		}
	}()

	return &msg, nil
}

func CloseServer(msg *ClientMsg)  {
	close(msg.MsgChan)
}
