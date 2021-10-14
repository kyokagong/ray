package runtime_test

import (
	"ray/internal/runtime"
	c "ray/internal/runtine/config"
	"testing"
)

func TestInitRuntime(t *testing.T) {
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
	runtime.InitRayRuntime(config)
}
