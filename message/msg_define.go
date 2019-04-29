// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.
package message

import (
	"github.com/pineal-niwan/busybox/binary"
	"github.com/pineal-niwan/busybox/fast_rpc"
)

//获取单键
type MsgReqGet struct {
	ReqGet
}

//获取命令行
func (msg *MsgReqGet) GetCmd() uint16 {
	return uint16(1)
}

//获取版本号
func (msg *MsgReqGet) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgReqGet) GetCode() uint32 {
	return uint32(1) | (uint32(0) << 16)
}

//序列化
func (msg *MsgReqGet) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteReqGet(msg.ReqGet)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(1),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgReqGet) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.ReqGet, err = reader.ReadReqGet()
	return err
}

//返回单键
type MsgRspGet struct {
	RspGet
}

//获取命令行
func (msg *MsgRspGet) GetCmd() uint16 {
	return uint16(2)
}

//获取版本号
func (msg *MsgRspGet) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgRspGet) GetCode() uint32 {
	return uint32(2) | (uint32(0) << 16)
}

//序列化
func (msg *MsgRspGet) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteRspGet(msg.RspGet)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(2),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgRspGet) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.RspGet, err = reader.ReadRspGet()
	return err
}

//获取多键
type MsgReqGetList struct {
	ReqGetList
}

//获取命令行
func (msg *MsgReqGetList) GetCmd() uint16 {
	return uint16(3)
}

//获取版本号
func (msg *MsgReqGetList) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgReqGetList) GetCode() uint32 {
	return uint32(3) | (uint32(0) << 16)
}

//序列化
func (msg *MsgReqGetList) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteReqGetList(msg.ReqGetList)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(3),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgReqGetList) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.ReqGetList, err = reader.ReadReqGetList()
	return err
}

//返回一组键值对
type MsgRspGetList struct {
	RspGetList
}

//获取命令行
func (msg *MsgRspGetList) GetCmd() uint16 {
	return uint16(4)
}

//获取版本号
func (msg *MsgRspGetList) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgRspGetList) GetCode() uint32 {
	return uint32(4) | (uint32(0) << 16)
}

//序列化
func (msg *MsgRspGetList) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteRspGetList(msg.RspGetList)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(4),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgRspGetList) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.RspGetList, err = reader.ReadRspGetList()
	return err
}

//设置键值对
type MsgReqSet struct {
	ReqSet
}

//获取命令行
func (msg *MsgReqSet) GetCmd() uint16 {
	return uint16(5)
}

//获取版本号
func (msg *MsgReqSet) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgReqSet) GetCode() uint32 {
	return uint32(5) | (uint32(0) << 16)
}

//序列化
func (msg *MsgReqSet) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteReqSet(msg.ReqSet)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(5),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgReqSet) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.ReqSet, err = reader.ReadReqSet()
	return err
}

//返回设置单键结果
type MsgRspSet struct {
	RspSet
}

//获取命令行
func (msg *MsgRspSet) GetCmd() uint16 {
	return uint16(6)
}

//获取版本号
func (msg *MsgRspSet) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgRspSet) GetCode() uint32 {
	return uint32(6) | (uint32(0) << 16)
}

//序列化
func (msg *MsgRspSet) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteRspSet(msg.RspSet)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(6),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgRspSet) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.RspSet, err = reader.ReadRspSet()
	return err
}

//设置多键
type MsgReqSetList struct {
	ReqSetList
}

//获取命令行
func (msg *MsgReqSetList) GetCmd() uint16 {
	return uint16(7)
}

//获取版本号
func (msg *MsgReqSetList) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgReqSetList) GetCode() uint32 {
	return uint32(7) | (uint32(0) << 16)
}

//序列化
func (msg *MsgReqSetList) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteReqSetList(msg.ReqSetList)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(7),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgReqSetList) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.ReqSetList, err = reader.ReadReqSetList()
	return err
}

//返回设置多键结果
type MsgRspSetList struct {
	RspSetList
}

//获取命令行
func (msg *MsgRspSetList) GetCmd() uint16 {
	return uint16(8)
}

//获取版本号
func (msg *MsgRspSetList) GetVersion() uint16 {
	return uint16(0)
}

//获取Code
func (msg *MsgRspSetList) GetCode() uint32 {
	return uint32(8) | (uint32(0) << 16)
}

//序列化
func (msg *MsgRspSetList) Marshal(buf []byte, option *binary.Option) (int, []byte, error) {
	writer, err := NewWriteIncHandlerWithOption(buf, option)
	if err != nil {
		return 0, nil, err
	}
	//先跳过消息头
	err = writer.MovePos(uint32(fast_rpc.MsgHeadSize))
	if err != nil {
		return 0, nil, err
	}
	//写消息内容
	err = writer.WriteRspSetList(msg.RspSetList)
	if err != nil {
		return 0, nil, err
	}
	size := writer.ResetPos(0)
	//回填消息头
	contentSize := uint32(size - fast_rpc.MsgHeadSize)
	err = fast_rpc.MarshalMsgHead(writer.BinaryHandler,
		fast_rpc.MsgHead{
			Size:    contentSize,
			Cmd:     uint16(8),
			Version: uint16(0),
		})
	if err != nil {
		return 0, nil, err
	}
	writer.ResetPos(size)
	return size, writer.Data(), err
}

//反序列化
func (msg *MsgRspSetList) Unmarshal(buf []byte, option *binary.Option) error {
	reader, err := NewReadIncHandlerWithOption(buf, option)
	if err != nil {
		return err
	}
	//读消息内容
	msg.RspSetList, err = reader.ReadRspSetList()
	return err
}
