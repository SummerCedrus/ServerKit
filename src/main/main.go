package main

import (
	"fmt"
	."netkit"
)

func main(){
	msg, err := NewServer("127.0.0.1:8080")

	if nil != err{
		fmt.Errorf("Create New Server Error [%s]", err.Error())
	}

	go msgHandle(msg)

	conn := ConnectServer("127.0.0.1:8080", "tcp4")

	conn.Write([]byte("Hello Server!"))
}

func msgHandle(msg *ClientMsg) {
	for {
		select {
		case b, ok := <-msg.MsgChan:

			if ok {
				fmt.Println(string(b))
			}
		//case cls, ok := <-msg.CloseChan:
		//	if ok && cls {
		//		close(msg.MsgChan)
		//	}
		}
	}
}
