// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.
package message

import (
	"context"
	"github.com/pineal-niwan/busybox/fast_rpc"
)

//传入key，获取value
func Get(
	ctx context.Context,
	cli *fast_rpc.Cli,
	input ReqGet,
	retryTimes int) (*RspGet, error) {

	//调用RPC - 发送消息后接收消息
	outMsg, err := cli.CallWithRetry(
		ctx,
		&MsgReqGet{
			ReqGet: input,
		},
		retryTimes)
	if err != nil {
		return nil, err
	}
	//检查是否是期望的消息
	output, ok := outMsg.(*MsgRspGet)
	if !ok {
		return nil, fast_rpc.ErrNotExpectMsg
	}
	return &output.RspGet, nil
}

//设置key-value
func Set(
	ctx context.Context,
	cli *fast_rpc.Cli,
	input ReqSet,
	retryTimes int) (*RspSet, error) {

	//调用RPC - 发送消息后接收消息
	outMsg, err := cli.CallWithRetry(
		ctx,
		&MsgReqSet{
			ReqSet: input,
		},
		retryTimes)
	if err != nil {
		return nil, err
	}
	//检查是否是期望的消息
	output, ok := outMsg.(*MsgRspSet)
	if !ok {
		return nil, fast_rpc.ErrNotExpectMsg
	}
	return &output.RspSet, nil
}

//获取一组key
func GetList(
	ctx context.Context,
	cli *fast_rpc.Cli,
	input ReqGetList,
	retryTimes int) (*RspGetList, error) {

	//调用RPC - 发送消息后接收消息
	outMsg, err := cli.CallWithRetry(
		ctx,
		&MsgReqGetList{
			ReqGetList: input,
		},
		retryTimes)
	if err != nil {
		return nil, err
	}
	//检查是否是期望的消息
	output, ok := outMsg.(*MsgRspGetList)
	if !ok {
		return nil, fast_rpc.ErrNotExpectMsg
	}
	return &output.RspGetList, nil
}

//设置一组key-value
func SetList(
	ctx context.Context,
	cli *fast_rpc.Cli,
	input ReqSetList,
	retryTimes int) (*RspSetList, error) {

	//调用RPC - 发送消息后接收消息
	outMsg, err := cli.CallWithRetry(
		ctx,
		&MsgReqSetList{
			ReqSetList: input,
		},
		retryTimes)
	if err != nil {
		return nil, err
	}
	//检查是否是期望的消息
	output, ok := outMsg.(*MsgRspSetList)
	if !ok {
		return nil, fast_rpc.ErrNotExpectMsg
	}
	return &output.RspSetList, nil
}
