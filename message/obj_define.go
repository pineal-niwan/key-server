// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.

package message

//键值对定义
type KV struct {
	//键
	Key string
	//值
	Value []byte
}

//键错误对定义
type KErr struct {
	//键
	Key string
	//错误信息
	Err string
}

//获取单键
type ReqGet struct {
	//键
	Key string
}

//返回单键
type RspGet struct {
	//字节流值
	Value []byte
	//获取过程是否有err,如果有,则此字符串表示error内容
	Err string
}

//获取多键
type ReqGetList struct {
	//一组键
	KeyList []string
}

//返回一组键值对
type RspGetList struct {
	//返回一组键值对
	KVList []KV
	//获取过程是否有err,如果有,则此字符串表示error内容
	Err string
}

//设置键值对
type ReqSet struct {
	//键值对
	KVPair KV
}

//返回设置单键结果
type RspSet struct {
	//设置过程是否有err,如果有,则此字符串表示error内容
	Err string
}

//设置多键
type ReqSetList struct {
	//一组键值对
	KVList []KV
}

//返回设置多键结果
type RspSetList struct {
	//设置过程是否有err,如果有,则此列表表示出错的键及其对应的原因
	KErrList []KErr
}
