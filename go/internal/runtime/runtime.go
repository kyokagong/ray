package runtime

import (
	"ray/internal/runtime/config"
	"ray/internal/runtime/object"
	"ray/internal/runtime/runner"
	"ray/internal/runtime/task"
)

type RunTime struct {
	functionManager task.FunctionManager
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
	Call(remoteFunctionHolder task.RemoteFunctionHolder, args interface{}) object.ObjectRef
	Put(interface{}) object.ObjectRef
	Get(object.ObjectRef) interface{}
	Wait()
}

// 初始化runtime
func InitRayRuntime(rayConfig config.RayConfig) {
	functiomManager := task.GetFunctionManager()
	if rayConfig.LocalMode {
		runtime = &RunTime{functiomManager, InitLocalRuntime()}
	} else {
		runner.StartRayHead(rayConfig)

	}
}
