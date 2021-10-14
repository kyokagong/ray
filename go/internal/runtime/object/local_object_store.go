package object

import (
	"log"
	"ray/internal/runtime/serializer"
	"time"

	"github.com/google/uuid"
	"github.com/muesli/cache2go"
)

var (
	LOCAL_CACHE_TABLE_NAME   = "rayGoCache"
	LOCAL_CACHE_DEFAULT_TIME = 3600 * 24 * time.Second
)

// LocalMode ObjectStore
type LocalObjectStore struct {
}

func CreateObjectRef() ObjectRef {
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	objRef := ObjectRef{u1.String()}
	return objRef
}

func (localObjectStore LocalObjectStore) Put(obj interface{}) ObjectRef {
	b := serializer.Serialize(obj)
	cache := cache2go.Cache(LOCAL_CACHE_TABLE_NAME)
	objRef := CreateObjectRef()
	cache.Add(objRef.GetID(), LOCAL_CACHE_DEFAULT_TIME, b)
	return objRef
}

func (localObjectStore LocalObjectStore) Get(objRef ObjectRef) interface{} {
	cache := cache2go.Cache(LOCAL_CACHE_TABLE_NAME)
	res, err := cache.Value(objRef.GetID())
	if err == nil {
		b := res.Data()
		return serializer.Deserialize(b.([]byte))
	} else {
		log.Fatal("objRef: ", objRef.ID, "|", err)
		return nil
	}
}

// Initialize LocalObjectStore
func InitLocalObjectStore() LocalObjectStore {
	return LocalObjectStore{}
}
