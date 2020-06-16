package protocol

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"errors"
)

func ReflectMessage(Cmd uint32) (proto.Message, error){
	switch Cmd {
	case CMD_GET_ITEM_INFO_REQ:
		return &ItemInfo{}, nil
	default:
		fmt.Printf("Error Cmd [%d]", Cmd)
		return nil, errors.New("Error Cmd")
	}
}
