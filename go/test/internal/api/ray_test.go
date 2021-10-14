package ray_test

import (
	"fmt"
	"ray/internal/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Foo(arr ...interface{}) interface{} {
	a := arr[0].(int)
	b := arr[1].(int)
	c := a + b
	return c
}

func TestInitRuntime(t *testing.T) {
	api.RayInitLocalMode()
	objRef := api.RayPut("test")
	fmt.Println("objRef: ", objRef)
	v := api.RayGet(objRef)
	assert := assert.New(t)
	assert.Equal(v, "test", "")
	api.RayTask(Foo)
}
