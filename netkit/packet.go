package netkit

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"github/SummerCedrus/ServerKit/protocol"
)

//消息打包约定结构
//|________|________________|_____________________
//	2 byte |  4 byte        |
//len(body)| cmd            |  body(data)

func Packet(msg *Message) ([]byte, error){
	//序列化msg.Msg
	data, err := encode(msg.Cmd, msg.Msg)
	if nil == data{
		return nil, err
	}
	netMsg := make([]byte, 0)
	//append head
	bodyLen := len(data)
	headBuff := make([]byte,2)
	headBuff[0] = byte(bodyLen >> 8)
	headBuff[1] = byte(bodyLen)
	fmt.Println("packet ",headBuff)
	netMsg = append(netMsg, headBuff...)
	//append cmd
	cmdBuff := make([]byte,4)
	cmdBuff[0] = byte(msg.Cmd >> 24)
	cmdBuff[1] = byte(msg.Cmd >> 16)
	cmdBuff[2] = byte(msg.Cmd >> 8)
	cmdBuff[3] = byte(msg.Cmd)
	fmt.Println("packet ",cmdBuff)
	netMsg = append(netMsg, cmdBuff...)
	//append body
	netMsg = append(netMsg, data...)

	return netMsg, nil
}

func UnPacket(netMsg []byte) (*Message, error){
	buff := netMsg[0:2]
	head := uint16(buff[0]) << 8 | uint16(buff[1])
	buff = netMsg[2:6]
	cmd := uint32(buff[0]) << 24 | uint32(buff[1]) << 16 | uint32(buff[2]) << 8 | uint32(buff[3])
	buff = netMsg[6:head+6]
	pbMsg, err := protocol.ReflectMessage(cmd)
	if nil == pbMsg{
		return nil, err
	}
	err = decode(cmd, buff, pbMsg)
	if nil != err{
		return nil, err
	}
	msg := &Message{
		Cmd:cmd,
		Msg:pbMsg,
	}

	return msg, nil
}

func encode(Cmd uint32, msg proto.Message) ([]byte, error){
	data, err := proto.Marshal(msg)
	if nil != err{
		fmt.Printf("Encode Message [%d] Error [%s]!!!", Cmd, err.Error())
		return nil, err
	}

	return data, nil
}

func decode(Cmd uint32,data []byte, msg proto.Message) error{
	err := proto.Unmarshal(data, msg)
	if nil != err{
		fmt.Printf("Decode Message [%d] Error!!!", Cmd)
		return err
	}

	return nil
}
