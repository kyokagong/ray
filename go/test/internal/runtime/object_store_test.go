package object_store_test

import (
	"fmt"
	"ray/internal/runtime/object"
	"testing"

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
