package serializer

import (
	"fmt"
	"log"

	"github.com/vmihailenco/msgpack"
)

func Serialize(obj interface{}) []byte {
	b, err := msgpack.Marshal(obj) // 将结构体转化为二进制流
	if err != nil {
		fmt.Printf("msgpack marshal failed,err:%v", err)
		log.Fatal(err)
	}
	return b
}

func Deserialize(b []byte) interface{} {
	var out interface{}
	err := msgpack.Unmarshal(b, &out)
	if err != nil {
		fmt.Printf("msgpack Unmarshal failed,err:%v", err)
		log.Fatal(err)
	}
	return out
}
