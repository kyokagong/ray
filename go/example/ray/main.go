package main

import (
	"fmt"

	"github.com/ray-project/ray/go/internal/api"
)

func Foo(a int64, b int64) int64 {
	c := a + b
	return c
}

func init() {
	api.RayRemote(Foo)
}

func main() {

	api.RayInitLocalMode()
	objRef, _ := api.RayPut("test")
	fmt.Println("objRef: ", objRef)
	v, _ := api.RayGet(objRef)
	fmt.Printf("first get value: %v", v)
	objRef, err := api.RayTask(Foo).Remote(1, 2)
	if err != nil {
		panic(fmt.Sprintf("Remote error: %v", err))
	}
	ret, _ := api.RayGet(objRef)
	fmt.Printf("remote ret: %v", ret)
}
