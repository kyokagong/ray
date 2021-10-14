package runtime

import (
	"ray/internal/runtime/object"
	"ray/internal/runtime/task"
)

// Local GO Ray Runtime, contains LocalObjectStore
type LocalRayRuntime struct {
	ObjectStore   object.LocalObjectStore
	TaskSubmitter task.TaskSubmitter
}

// Submit Task to Task Executor
func (localRayRuntime LocalRayRuntime) Call(remoteFunctionHolder task.RemoteFunctionHolder, args interface{}) object.ObjectRef {
	return object.ObjectRef{"test"}
}

func (localRayRuntime LocalRayRuntime) Put(obj interface{}) object.ObjectRef {
	return localRayRuntime.ObjectStore.Put(obj)
}

func (localRayRuntime LocalRayRuntime) Get(objRef object.ObjectRef) interface{} {
	return localRayRuntime.ObjectStore.Get(objRef)
}

func (localRayRuntime LocalRayRuntime) Wait() {

}

func InitLocalRuntime() LocalRayRuntime {
	return LocalRayRuntime{
		object.InitLocalObjectStore(),
		task.NewLocalTaskSubmitter()}
}
