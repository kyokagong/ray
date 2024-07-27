package runtime

import (
	"github.com/ray-project/ray/go/internal/runtime/config"
	"github.com/ray-project/ray/go/internal/runtime/object"
	"github.com/ray-project/ray/go/internal/runtime/runner"
	"github.com/ray-project/ray/go/internal/runtime/task"
)

type RunTime struct {
	functionManager *task.FunctionManager
	Instance        RayRuntimeInterface
}

// Get Runtime Instance
func (runtime *RunTime) GetInstance() RayRuntimeInterface {
	return runtime.Instance
}

// A static runtime
var runtime *RunTime

// Return singleton runtime
func GetRuntime() *RunTime {
	return runtime
}

// RayRuntime contains basic runtime function
type RayRuntimeInterface interface {
	Call(*task.RemoteFunctionHolder, []interface{}) (object.ObjectRef, error)
	Put(interface{}) (object.ObjectRef, error)
	Get(object.ObjectRef) (interface{}, error)
	Wait() error
}

// 初始化runtime
func InitRayRuntime(rayConfig config.RayConfig) {
	functiomManager := task.GetFunctionManager()
	if rayConfig.LocalMode {
		runtimeInstance := NewLocalRuntime()
		runtime = &RunTime{functiomManager, &runtimeInstance}
	} else {
		runner.StartRayHead(rayConfig)
	}
}
