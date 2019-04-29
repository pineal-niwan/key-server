// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.

package message

import (
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
)

var (
	InitMsgParseHandlerHash map[uint32]fast_rpc.MsgParseHandler
)

func init() {
	var key uint32
	InitMsgParseHandlerHash = make(map[uint32]fast_rpc.MsgParseHandler)
	//生成ReqGet解析函数 command:1 version:0
	key = uint32(1) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgReqGet{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成RspGet解析函数 command:2 version:0
	key = uint32(2) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgRspGet{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成ReqGetList解析函数 command:3 version:0
	key = uint32(3) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgReqGetList{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成RspGetList解析函数 command:4 version:0
	key = uint32(4) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgRspGetList{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成ReqSet解析函数 command:5 version:0
	key = uint32(5) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgReqSet{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成RspSet解析函数 command:6 version:0
	key = uint32(6) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgRspSet{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成ReqSetList解析函数 command:7 version:0
	key = uint32(7) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgReqSetList{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
	//生成RspSetList解析函数 command:8 version:0
	key = uint32(8) | (uint32(0) << 16)
	InitMsgParseHandlerHash[key] = func(data []byte, option *binary.Option) (fast_rpc.IMsg, error) {
		msg := &MsgRspSetList{}
		err := msg.Unmarshal(data, option)
		return msg, err
	}
}
