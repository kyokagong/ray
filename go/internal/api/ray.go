package api

import (
	"github.com/ray-project/ray/go/internal/runtime"
	c "github.com/ray-project/ray/go/internal/runtime/config"
	"github.com/ray-project/ray/go/internal/runtime/object"
	"github.com/ray-project/ray/go/internal/runtime/task"
)

func RayInitLocalMode() {
	var redisAddress string
	var redisPort int
	var objectManagerPort int
	var nodeManagerPort int
	var gcsServerPort int
	nodeIpAddress := "0.0.0.0"
	var rayletIpAddress string
	var rayClientServerPort int
	var redisPassword string
	localMode := true
	config := c.CreateRayConfig(
		redisAddress,
		redisPort,
		objectManagerPort,
		nodeManagerPort,
		gcsServerPort,
		nodeIpAddress,
		rayletIpAddress,
		rayClientServerPort,
		redisPassword,
		localMode)
	RayInit(config)
}

// Init Ray
func RayInit(config c.RayConfig) {
	runtime.InitRayRuntime(config)
}

// Create a ray remote task
func RayTask(fn interface{}) TaskCaller {
	return NewTaskCaller(task.NewFunctionWrapper(fn))
}

func RayPut(obj interface{}) (object.ObjectRef, error) {
	return runtime.GetRuntime().GetInstance().Put(obj)
}

func RayGet(objRef object.ObjectRef) (interface{}, error) {
	return runtime.GetRuntime().GetInstance().Get(objRef)
}

func RayRemote(fn interface{}) {
	funcWrapper := task.NewFunctionWrapper(fn)
	task.GetFunctionManager().RegisterRemoteFunction(funcWrapper)
}

func RayShutdown() {

}
