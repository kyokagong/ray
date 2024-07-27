package object

import (
	"log"
	"time"

	"github.com/ray-project/ray/go/internal/logger"
	"github.com/ray-project/ray/go/internal/runtime/serializer"

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

func NewObjectRef() ObjectRef {
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	objRef := ObjectRef{u1.String()}
	return objRef
}

func (store *LocalObjectStore) Put(obj interface{}) (ObjectRef, error) {
	objRef := NewObjectRef()
	err := store.PutRaw(objRef, obj)
	if err != nil {
		return objRef, err
	}
	return objRef, nil
}

func (store *LocalObjectStore) Get(objRef ObjectRef) (interface{}, error) {
	cache := cache2go.Cache(LOCAL_CACHE_TABLE_NAME)
	res, err := cache.Value(objRef.GetID())
	if err == nil {
		b := res.Data()
		value, err := serializer.Deserialize(b.([]byte))
		if err != nil {
			return nil, err
		}
		return value, nil
	} else {
		logger.Errorf("objRef: %v, error: %v", objRef.ID, err)
		return nil, err
	}
}

func (store *LocalObjectStore) PutRaw(objRef ObjectRef, obj interface{}) error {
	b, err := serializer.Serialize(obj)
	if err != nil {
		return err
	}
	cache := cache2go.Cache(LOCAL_CACHE_TABLE_NAME)
	cache.Add(objRef.GetID(), LOCAL_CACHE_DEFAULT_TIME, b)
	return nil
}

// Initialize LocalObjectStore
func NewLocalObjectStore() LocalObjectStore {
	return LocalObjectStore{}
}
