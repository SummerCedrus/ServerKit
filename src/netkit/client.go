package netkit

import (
	"net"
	"fmt"
)

func ConnectServer(addr string, netWork string) *net.TCPConn{
	tcpAddr, err := net.ResolveTCPAddr(netWork, addr)

	if nil != err{
		return nil
	}

	conn, err := net.DialTCP(netWork, nil, tcpAddr)
	if nil != err{
		fmt.Errorf("DialTcp Error [%s]", err.Error())
		return nil
	}
	//conn.Write([]byte("Hello Server!"))

	return conn
}
