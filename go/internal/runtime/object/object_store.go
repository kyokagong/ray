package object

// ObjectStore Interface
type ObjectStoreInterface interface {
	Put(ObjectRef, interface{})
	Get(ObjectRef) interface{}
}
