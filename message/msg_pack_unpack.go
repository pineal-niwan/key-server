// Code generation
// !!! Do not edit it.
// !!! Use code gen tool to generate.

package message

import (
	"github.com/pineal-niwan/busybox/binary"
)

//处理key-value
type IncHandler struct {
	*binary.BinaryHandler
}

//反序列化handler,读取字节流到对象中
func NewReadIncHandlerWithOption(data []byte, option *binary.Option) (*IncHandler, error) {
	binHandler, err := binary.NewReadBinaryHandler(data, option)
	if err != nil {
		return nil, err
	} else {
		return &IncHandler{
			BinaryHandler: binHandler,
		}, nil
	}
}

//序列化handler,将对象转化成字节流
func NewWriteIncHandlerWithOption(data []byte, option *binary.Option) (*IncHandler, error) {
	binHandler, err := binary.NewWriteBinaryHandler(data, option)
	if err != nil {
		return nil, err
	} else {
		return &IncHandler{
			BinaryHandler: binHandler,
		}, nil
	}
}

//读取KV
func (p *IncHandler) ReadKV() (ret KV, err error) {
	ret.Key, err = p.ReadString()
	if err != nil {
		return
	}
	ret.Value, err = p.ReadByteArray()
	if err != nil {
		return
	}
	return
}

//写入KV
func (p *IncHandler) WriteKV(v KV) (err error) {
	err = p.WriteString(v.Key)
	if err != nil {
		return
	}
	err = p.WriteByteArray(v.Value)
	if err != nil {
		return
	}
	return
}

//读取KV数组
func (p *IncHandler) ReadKVArray() (ret []KV, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]KV, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadKV()
		if err != nil {
			return
		}
	}
	return
}

//写入KV数组
func (p *IncHandler) WriteKVArray(v []KV) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteKV(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取KErr
func (p *IncHandler) ReadKErr() (ret KErr, err error) {
	ret.Key, err = p.ReadString()
	if err != nil {
		return
	}
	ret.Err, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入KErr
func (p *IncHandler) WriteKErr(v KErr) (err error) {
	err = p.WriteString(v.Key)
	if err != nil {
		return
	}
	err = p.WriteString(v.Err)
	if err != nil {
		return
	}
	return
}

//读取KErr数组
func (p *IncHandler) ReadKErrArray() (ret []KErr, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]KErr, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadKErr()
		if err != nil {
			return
		}
	}
	return
}

//写入KErr数组
func (p *IncHandler) WriteKErrArray(v []KErr) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteKErr(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取ReqGet
func (p *IncHandler) ReadReqGet() (ret ReqGet, err error) {
	ret.Key, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入ReqGet
func (p *IncHandler) WriteReqGet(v ReqGet) (err error) {
	err = p.WriteString(v.Key)
	if err != nil {
		return
	}
	return
}

//读取ReqGet数组
func (p *IncHandler) ReadReqGetArray() (ret []ReqGet, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]ReqGet, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadReqGet()
		if err != nil {
			return
		}
	}
	return
}

//写入ReqGet数组
func (p *IncHandler) WriteReqGetArray(v []ReqGet) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteReqGet(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取RspGet
func (p *IncHandler) ReadRspGet() (ret RspGet, err error) {
	ret.Value, err = p.ReadByteArray()
	if err != nil {
		return
	}
	ret.Err, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入RspGet
func (p *IncHandler) WriteRspGet(v RspGet) (err error) {
	err = p.WriteByteArray(v.Value)
	if err != nil {
		return
	}
	err = p.WriteString(v.Err)
	if err != nil {
		return
	}
	return
}

//读取RspGet数组
func (p *IncHandler) ReadRspGetArray() (ret []RspGet, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]RspGet, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadRspGet()
		if err != nil {
			return
		}
	}
	return
}

//写入RspGet数组
func (p *IncHandler) WriteRspGetArray(v []RspGet) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteRspGet(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取ReqGetList
func (p *IncHandler) ReadReqGetList() (ret ReqGetList, err error) {
	ret.KeyList, err = p.ReadStringArray()
	if err != nil {
		return
	}
	return
}

//写入ReqGetList
func (p *IncHandler) WriteReqGetList(v ReqGetList) (err error) {
	err = p.WriteStringArray(v.KeyList)
	if err != nil {
		return
	}
	return
}

//读取ReqGetList数组
func (p *IncHandler) ReadReqGetListArray() (ret []ReqGetList, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]ReqGetList, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadReqGetList()
		if err != nil {
			return
		}
	}
	return
}

//写入ReqGetList数组
func (p *IncHandler) WriteReqGetListArray(v []ReqGetList) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteReqGetList(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取RspGetList
func (p *IncHandler) ReadRspGetList() (ret RspGetList, err error) {
	ret.KVList, err = p.ReadKVArray()
	if err != nil {
		return
	}
	ret.Err, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入RspGetList
func (p *IncHandler) WriteRspGetList(v RspGetList) (err error) {
	err = p.WriteKVArray(v.KVList)
	if err != nil {
		return
	}
	err = p.WriteString(v.Err)
	if err != nil {
		return
	}
	return
}

//读取RspGetList数组
func (p *IncHandler) ReadRspGetListArray() (ret []RspGetList, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]RspGetList, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadRspGetList()
		if err != nil {
			return
		}
	}
	return
}

//写入RspGetList数组
func (p *IncHandler) WriteRspGetListArray(v []RspGetList) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteRspGetList(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取ReqSet
func (p *IncHandler) ReadReqSet() (ret ReqSet, err error) {
	ret.KVPair, err = p.ReadKV()
	if err != nil {
		return
	}
	return
}

//写入ReqSet
func (p *IncHandler) WriteReqSet(v ReqSet) (err error) {
	err = p.WriteKV(v.KVPair)
	if err != nil {
		return
	}
	return
}

//读取ReqSet数组
func (p *IncHandler) ReadReqSetArray() (ret []ReqSet, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]ReqSet, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadReqSet()
		if err != nil {
			return
		}
	}
	return
}

//写入ReqSet数组
func (p *IncHandler) WriteReqSetArray(v []ReqSet) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteReqSet(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取RspSet
func (p *IncHandler) ReadRspSet() (ret RspSet, err error) {
	ret.Err, err = p.ReadString()
	if err != nil {
		return
	}
	return
}

//写入RspSet
func (p *IncHandler) WriteRspSet(v RspSet) (err error) {
	err = p.WriteString(v.Err)
	if err != nil {
		return
	}
	return
}

//读取RspSet数组
func (p *IncHandler) ReadRspSetArray() (ret []RspSet, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]RspSet, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadRspSet()
		if err != nil {
			return
		}
	}
	return
}

//写入RspSet数组
func (p *IncHandler) WriteRspSetArray(v []RspSet) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteRspSet(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取ReqSetList
func (p *IncHandler) ReadReqSetList() (ret ReqSetList, err error) {
	ret.KVList, err = p.ReadKVArray()
	if err != nil {
		return
	}
	return
}

//写入ReqSetList
func (p *IncHandler) WriteReqSetList(v ReqSetList) (err error) {
	err = p.WriteKVArray(v.KVList)
	if err != nil {
		return
	}
	return
}

//读取ReqSetList数组
func (p *IncHandler) ReadReqSetListArray() (ret []ReqSetList, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]ReqSetList, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadReqSetList()
		if err != nil {
			return
		}
	}
	return
}

//写入ReqSetList数组
func (p *IncHandler) WriteReqSetListArray(v []ReqSetList) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteReqSetList(v[i])
		if err != nil {
			return
		}
	}
	return
}

//读取RspSetList
func (p *IncHandler) ReadRspSetList() (ret RspSetList, err error) {
	ret.KErrList, err = p.ReadKErrArray()
	if err != nil {
		return
	}
	return
}

//写入RspSetList
func (p *IncHandler) WriteRspSetList(v RspSetList) (err error) {
	err = p.WriteKErrArray(v.KErrList)
	if err != nil {
		return
	}
	return
}

//读取RspSetList数组
func (p *IncHandler) ReadRspSetListArray() (ret []RspSetList, err error) {
	var size uint32

	//读长度
	size, err = p.ReadArrayLen()
	if err != nil {
		return
	}
	//读内容
	ret = make([]RspSetList, size)
	for i := uint32(0); i < size; i++ {
		ret[i], err = p.ReadRspSetList()
		if err != nil {
			return
		}
	}
	return
}

//写入RspSetList数组
func (p *IncHandler) WriteRspSetListArray(v []RspSetList) (err error) {
	//写长度
	var size int
	if v == nil {
		size = 0
	} else {
		size = len(v)
	}
	err = p.WriteArrayLen(size)
	if err != nil {
		return
	}

	//写内容
	for i := 0; i < size; i++ {
		err = p.WriteRspSetList(v[i])
		if err != nil {
			return
		}
	}
	return
}
