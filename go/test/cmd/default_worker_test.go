package cmd_test

import (
	"testing"

	w "github.com/ray-project/ray/go/internal/runtime/worker"
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
