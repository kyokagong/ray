package api

import (
	"ray/internal/runtime"
	c "ray/internal/runtime/config"
	"ray/internal/runtime/object"
	"ray/internal/runtime/task"
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
func RayTask(function task.FuncType1) TaskCaller {
	return CreateTaskCaller(function)
}

func RayPut(obj interface{}) object.ObjectRef {
	return runtime.GetRuntime().GetInstance().Put(obj)
}

func RayGet(objRef object.ObjectRef) interface{} {
	return runtime.GetRuntime().GetInstance().Get(objRef)
}

func RayShutdown() {

}
