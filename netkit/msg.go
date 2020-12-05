package netkit

import (
	"net"
	"fmt"
	"strings"
	"time"
)

const (
	TCP_TIME_OUT = 60
)
func NewReciever(conn *net.TCPConn) {
	MaxSessionID++
	receiver := &Receiver{
		ID:        MaxSessionID,
		Conn:      conn,
		//CloseChan: make(chan bool),
		MsgChan: make(chan *Message, MaxMsgNum),
	}
	vIp := strings.Split(conn.RemoteAddr().String(), ":")
	ip := net.ParseIP(vIp[0])

	receiver.IP = ip
	mgr.ConnectChan <- receiver
	mgr.ReceiverPool[MaxSessionID] = receiver
	//所有session的消息都推入mgr的消息队列处理
	receiver.MsgChan = mgr.MsgChan
	receiver.run()
}
func NewSender(conn *net.TCPConn) *Sender{
	MaxSessionID++
	sender := &Sender{
		Queue: make(chan *Message, MaxMsgNum),
		Conn:conn,
		CloseChan:make(chan bool),
	}

	sender.run()

	return sender
}
//send data
func Send(conn *net.TCPConn, msg *Message) (int, error){
	data, err := Packet(msg)
	if nil != err{
		return 0, err
	}
	fmt.Println(data)
	data, err = EncryptDES_CBC(data)
	if nil != err{
		fmt.Errorf("encrypt failed![%s]", err.Error())
	}

	fmt.Println("Encrypted", data)


	return conn.Write(data)
}

func (mgr *ConnectMgr)GetReceiver(id int32) *Receiver{
	return mgr.ReceiverPool[id]
}

func (mgr *ConnectMgr)DelReceiver(id int32){
	delete(mgr.ReceiverPool, id)
}
func (sender *Sender)Send(msg *Message){
	sender.Queue <- msg
}

func (sender *Sender) run() {
	go func() {
		for {
			select {
			case msg, ok := <-sender.Queue:
				if ok {
					cnt, err := Send(sender.Conn, msg)
					if nil != err{
						fmt.Printf("cnt[%d] %s", cnt, err.Error())
						sender.Conn = ReConnectServer(sender.Conn)
					}
				}
			case _, ok := <-sender.CloseChan:
				if ok {
					return
				}
			}
		}
	}()
}

func (receiver *Receiver) readMsg() error {
	conn := receiver.Conn
	conn.SetDeadline(time.Now().Add(TCP_TIME_OUT*time.Second))
	buff := make([]byte, 128)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Read failed: ", err, conn)
		return err
	}
	fmt.Println(buff)
	data, err := DecryptDES_CBC(buff)
	fmt.Println("Decrypted ",data)
	msg, err := UnPacket(data)
	if nil != err {
		fmt.Errorf("Decrypt failed![%s]", err.Error())
		return err
	}
	receiver.MsgChan <- msg

	return nil
}


func (receiver *Receiver) run() {
	defer func() {
		receiver.Conn.Close()
		mgr.DelReceiver(receiver.ID)
		fmt.Println("close connect...")
	}()
	for {
		err := receiver.readMsg()
		if err != nil {
			fmt.Println("read msg err:",err.Error())
			return
		}
	}
}
