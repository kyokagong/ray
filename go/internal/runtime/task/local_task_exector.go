package task

import (
	"github.com/ray-project/ray/go/internal/logger"
	"github.com/ray-project/ray/go/internal/runtime/object"
)

type LocalTaskExecutor struct {
	ObjectStore *object.LocalObjectStore
}

type Result struct {
	Value interface{}
}

func (r *Result) IsEmpty() bool {
	return r.Value == nil
}

func (executor *LocalTaskExecutor) ExecuteTask(funcName string, args []interface{}) Result {
	funcManager := GetFunctionManager()

	rayFunction, err := funcManager.GetRayFunction(funcName)
	if err != nil {
		logger.Errorf("LocalTaskExecutor GetRayFunction err: %v", err)
		return Result{}
	}
	ret, err := rayFunction.Call(args...)
	if err != nil {
		logger.Errorf("LocalTaskExecutor ExecuteTask err: %v", err)
		return Result{}
	}
	return Result{
		Value: ret,
	}
}

func NewLocaLocalTaskExecutor(objectStore *object.LocalObjectStore) *LocalTaskExecutor {
	return &LocalTaskExecutor{
		objectStore,
	}
}
