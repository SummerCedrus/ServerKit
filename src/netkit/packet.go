package netkit

//消息打包约定结构
//|________|________________|_____________________
//	2 byte |  4 byte        |
//len(body)| cmd            |  body

func Packet(msg *Message) []byte{
	data := make([]byte,0)
	//append head
	bodyLen := len(msg.Data)
	headBuff := make([]byte,2)
	headBuff[0] = byte(bodyLen >> 8)
	headBuff[1] = byte(bodyLen)
	data = append(data, headBuff...)
	//append cmd
	cmdBuff := make([]byte,4)
	cmdBuff[0] = byte(msg.Cmd >> 24)
	cmdBuff[1] = byte(msg.Cmd >> 16)
	cmdBuff[2] = byte(msg.Cmd >> 8)
	cmdBuff[3] = byte(msg.Cmd)
	data = append(data, cmdBuff...)
	//apped body
	data = append(data, msg.Data...)

	return  data
}

func UnPacket(data []byte) *Message{
	buff := data[0:2]
	head := uint16(buff[0]) << 8 | uint16(buff[1])
	buff = data[2:6]
	cmd := uint32(buff[0]) << 24 | uint32(buff[1]) << 16 | uint32(buff[2]) << 8 | uint32(buff[3])
	buff = data[6:head+6]
	msg := &Message{
		Cmd:int32(cmd),
		Data:buff,
	}

	return msg
}