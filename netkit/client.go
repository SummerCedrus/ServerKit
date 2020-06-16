package netkit

import (
	"net"
	"fmt"
	"time"
)

func NewClient(addr string, netWork string) *Sender{
	conn := ConnectServer("127.0.0.1:8080", "tcp4")
	if nil == conn{
		fmt.Errorf("err connect server[%s]",addr)
		return nil
	}
	fmt.Println(conn.RemoteAddr(),conn.LocalAddr())
	sender := NewSender(conn)

	return sender
}

func ConnectServer(addr string, netWork string) *net.TCPConn{
	tcpAddr, err := net.ResolveTCPAddr(netWork, addr)

	if nil != err{
		return nil
	}
	var conn *net.TCPConn
	for {
		conn, err = net.DialTCP(netWork, nil, tcpAddr)
		if nil == err && nil != conn{
			break
		}
		fmt.Printf("DialTcp Error [%s] Reconnect...\n", err.Error())
		time.Sleep(time.Second)
	}

	//conn.Write([]byte("Hello Server!"))
	fmt.Printf("Connect [%s] success!", addr)
	return conn
}

func ReConnectServer(conn *net.TCPConn) *net.TCPConn{
	addr := conn.RemoteAddr()
	addr_str := addr.String()
	conn.Close()
	return ConnectServer(addr_str, "tcp4")
}
