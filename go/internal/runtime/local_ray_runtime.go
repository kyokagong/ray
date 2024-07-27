package runtime

import (
	"time"

	"github.com/ray-project/ray/go/internal/runtime/object"
	"github.com/ray-project/ray/go/internal/runtime/task"
)

// Local GO Ray Runtime, contains LocalObjectStore
type LocalRayRuntime struct {
	ObjectStore   *object.LocalObjectStore
	TaskSubmitter *task.LocalTaskSubmitter
}

// Submit Task to Task Executor
func (runtime *LocalRayRuntime) Call(remoteFunctionHolder *task.RemoteFunctionHolder, args []interface{}) (object.ObjectRef, error) {
	objRef, err := runtime.TaskSubmitter.SubmitTask(remoteFunctionHolder, args)
	if err != nil {
		return object.ObjectRef{}, err
	}
	return objRef, nil
}

func (runtime *LocalRayRuntime) Put(obj interface{}) (object.ObjectRef, error) {
	return runtime.ObjectStore.Put(obj)
}

func (runtime *LocalRayRuntime) Get(objRef object.ObjectRef) (interface{}, error) {
	for {
		if value, err := runtime.ObjectStore.Get(objRef); err == nil {
			return value, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (runtime *LocalRayRuntime) Wait() error {
	return nil
}

func NewLocalRuntime() LocalRayRuntime {
	localObjectStore := object.NewLocalObjectStore()
	localTaskSubmitter := task.NewLocalTaskSubmitter(&localObjectStore)
	return LocalRayRuntime{
		ObjectStore:   &localObjectStore,
		TaskSubmitter: &localTaskSubmitter,
	}
}
