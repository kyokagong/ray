package object

// ObjectRefInterface contains ObjectRef Basic operations
type ObjectRefInterface interface {
	GetID() string
	Equals(otherObjRef ObjectRef) bool
}

// ObjectRef is func Remote return
type ObjectRef struct {
	ID string
}

func (objRef ObjectRef) GetID() string {
	return objRef.ID
}

func (objRef ObjectRef) Equals(otherObjRef ObjectRef) bool {
	return objRef.ID == otherObjRef.ID
}
