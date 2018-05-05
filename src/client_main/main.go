package main

import (
	."netkit"
	"time"
	"fmt"
)

func main(){
	conn := ConnectServer("127.0.0.1:8080", "tcp4")
	conn.Write([]byte("Hello Server!"))
	fmt.Println("Send 'Hello Server!'")
	for {
		time.Sleep(time.Second)
	}
}
