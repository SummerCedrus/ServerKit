package main

import (
	."github/SummerCedrus/ServerKit/netkit"
	"time"
	"fmt"
	."github/SummerCedrus/ServerKit/protocol"
)

func main(){

	session := NewClient("127.0.0.1:8080", "tcp4")

	if nil == session{
		fmt.Errorf("connect server [%s] failed", "127.0.0.1:8080")
		return
	}

	for {
		info := ItemInfo{
			ID:1,
			Type:2,
			Name:"sword_1",
			Amount:1,
		}

		msg := &Message{
			Cmd:CMD_GET_ITEM_INFO_REQ,
			Msg:&info,
		}

		session.Send(msg)

		time.Sleep(10*time.Second)
	}
}

