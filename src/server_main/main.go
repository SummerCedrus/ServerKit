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
	mainWork(msg)
}

func msgHandle(msg *ClientMsg) (err interface{}){
	for {
		select {
		case b, ok := <-msg.MsgChan:

			if ok {
				fmt.Println(string(b))
			}
		case cls, ok := <-msg.CloseChan:
			if ok && cls {
				close(msg.MsgChan)
			}

			return 0
		}
	}
}

func mainWork(msg *ClientMsg){
	for retryCnt := 0; retryCnt < MAX_CONNECT_RETRY_TIME;{
		err := msgHandle(msg)
		if nil == err{
			break
		}else{
			retryCnt ++
		}
	}

	ShowDown()
}

func ShowDown(){
	fmt.Println("Server ShutDown!")
}
