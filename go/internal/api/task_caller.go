package api

import (
	"github.com/ray-project/ray/go/internal/runtime/task"

	"github.com/ray-project/ray/go/internal/runtime/object"

	r "github.com/ray-project/ray/go/internal/runtime"
)

type TaskCallerInterface interface {
	Remote() object.ObjectRef
}

type TaskCaller struct {
	runtime     r.RunTime
	funcWrapper task.FunctionWrapper
}

// Create Golang Task Caller
func NewTaskCaller(funcWrapper task.FunctionWrapper) TaskCaller {
	return TaskCaller{*r.GetRuntime(), funcWrapper}
}

// Submit ray function
func (taskCaller TaskCaller) Remote(args ...interface{}) (object.ObjectRef, error) {
	remoteFuncHolder := task.BuildRemoteFunctionHolder(taskCaller.funcWrapper)
	return taskCaller.runtime.GetInstance().Call(&remoteFuncHolder, args)
}
