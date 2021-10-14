package cmd_test

import (
	w "ray/internal/runtime/worker"
	"testing"
)

func TestStartDefaultWorker(t *testing.T) {
	var redisAddress string
	var redisPort int
	var objectManagerPort int
	var odeManagerPort int
	var gcsServerPort int
	nodeIpAddress := "0.0.0.0"
	var rayletIpAddress string
	var rayClientServerPort int
	var redisPassword string
	localMode := true
	w.StartDefaultWorker(
		redisAddress,
		redisPort,
		objectManagerPort,
		odeManagerPort,
		gcsServerPort,
		nodeIpAddress,
		rayletIpAddress,
		rayClientServerPort,
		redisPassword,
		localMode,
	)
}
