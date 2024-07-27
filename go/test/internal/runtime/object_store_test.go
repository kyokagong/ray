package object_store_test

import (
	"fmt"
	"testing"

	"github.com/ray-project/ray/go/internal/runtime/object"

	"github.com/stretchr/testify/assert"
)

func TestLocalObjectStore(t *testing.T) {
	objectStore := object.LocalObjectStore{}
	objRef := objectStore.Put("test")
	v := objectStore.Get(objRef)
	fmt.Println("TestLocalObjectStore:", v)
	assert := assert.New(t)
	assert.Equal(v, "test", "")
}
