package worker

import (
	"fmt"
	api "ray/internal/api"
	c "ray/internal/runtime/config"
)

/*
#include "go/internal/runtime/worker/cgo_core_worker.h"
*/
import "C"

func StartDefaultWorker(
	redisAddress string,
	redisPort int,
	objectManagerPort int,
	nodeManagerPort int,
	gcsServerPort int,
	nodeIpAddress string,
	rayletIpAddress string,
	rayClientServerPort int,
	redisPassword string,
	localMode bool) {

	config := c.CreateRayConfig(
		redisAddress, redisPort, objectManagerPort, nodeManagerPort, gcsServerPort, nodeIpAddress, rayletIpAddress, rayClientServerPort, redisPassword, localMode)
	api.RayInit(config)
	fmt.Println("start ray golang worker")

	isLocalMode := 0
	if localMode {
		isLocalMode = 1
	}
	C.CCoreWorkerRunTaskExecutionLoop(
		C.CString(config.NodeIpAddress),
		C.int(config.NodeManagerPort),
		C.int(1),
		C.CString(config.RedisAddress),
		C.CString(config.RedisPassword),
		C.int(config.RedisPort),
		C.CString(config.RayletIpAddress),
		C.int(isLocalMode),
	)
}
