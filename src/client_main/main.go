package main

import (
	."netkit"
	"time"
	"fmt"
)

func main(){

	session := NewClient("127.0.0.1:8080", "tcp4")

	if nil == session{
		fmt.Errorf("connect server [%s] failed", "127.0.0.1:8080")
		return
	}

	for {
		session.Send([]byte("Hello Server!!!"))

		time.Sleep(20*time.Second)
	}
}

