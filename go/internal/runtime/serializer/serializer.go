package serializer

import (
	"github.com/ray-project/ray/go/internal/logger"
	"github.com/vmihailenco/msgpack"
)

func Serialize(obj interface{}) ([]byte, error) {
	b, err := msgpack.Marshal(obj) // 将结构体转化为二进制流
	if err != nil {
		logger.Errorf("msgpack marshal failed,err: %v", err)
		return []byte{}, err
	}
	return b, nil
}

func Deserialize(b []byte) (interface{}, error) {
	var out interface{}
	err := msgpack.Unmarshal(b, &out)
	if err != nil {
		logger.Errorf("msgpack Unmarshal failed,err:%v", err)
		return nil, err
	}
	return out, nil
}
