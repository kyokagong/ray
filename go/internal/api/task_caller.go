package api

import (
	r "ray/internal/runtime"
	"ray/internal/runtime/object"
	"ray/internal/runtime/task"
)

type TaskCallerInterface interface {
	Remote() object.ObjectRef
}

type TaskCaller struct {
	runtime  r.RunTime
	function task.FuncType1
}

// Create Golang Task Caller
func CreateTaskCaller(function task.FuncType1) TaskCaller {
	return TaskCaller{*r.GetRuntime(), function}
}

// Submit ray function
func (taskCaller TaskCaller) Remote() object.ObjectRef {
	return object.ObjectRef{"test"}
}
