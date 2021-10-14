package config

import "strconv"

type RayConfig struct {
	RedisAddress        string
	RedisPort           int
	ObjectManagerPort   int
	NodeManagerPort     int
	GcsServerPort       int
	NodeIpAddress       string
	RayletIpAddress     string
	RayClientServerPort int
	RedisPassword       string
	LocalMode           bool
	RayOpts             rayConfigOptions
}

func (rayConfig *RayConfig) GetHeadArgs() []string {
	//--head --node-ip-address=$MY_POD_IP --port=6379 --redis-shard-ports=6380,6381 --num-cpus=$MY_CPU_REQUEST --object-manager-port=12345 --node-manager-port=12346
	args := []string{}
	if rayConfig.NodeIpAddress != "" {
		args = append(args, "--node-ip-address", rayConfig.NodeIpAddress)
	}
	if rayConfig.RedisPort != 0 {
		args = append(args, "--port", strconv.Itoa(rayConfig.RedisPort))
	}
	if rayConfig.RayOpts.NumCpus != 0 {
		args = append(args, "--num-cpus", strconv.Itoa(rayConfig.RayOpts.NumCpus))
	}
	if rayConfig.ObjectManagerPort != 0 {
		args = append(args, "--object-manager-port", strconv.Itoa(rayConfig.ObjectManagerPort))
	}
	if rayConfig.NodeManagerPort != 0 {
		args = append(args, "--node-manager-port", strconv.Itoa(rayConfig.NodeManagerPort))
	}
	return args
}

type rayConfigOptions struct {
	NumCpus               int
	NumGpus               int
	PlasmaDirectory       string
	PlasmaStoreSocketName string
	RayletSocketName      string
}

type RayConfigOptions interface {
	apply(*rayConfigOptions)
}

var DefaultRayConfigOptions = rayConfigOptions{
	NumCpus:               1,
	NumGpus:               0,
	PlasmaDirectory:       "/tmp/ray/session/plasma",
	PlasmaStoreSocketName: "/tmp/ray/session/plasma",
	RayletSocketName:      "/tmp/ray/session/plasma",
}

type tempFunc func(*rayConfigOptions)

type funcRayConfigOptions struct {
	f tempFunc
}

func (frayOptions *funcRayConfigOptions) apply(c *rayConfigOptions) {
	frayOptions.f(c)
}

func newFuncRayConfigOptions(f tempFunc) *funcRayConfigOptions {
	return &funcRayConfigOptions{f: f}
}

func AddNumCpus(NumCpus int) RayConfigOptions {
	return newFuncRayConfigOptions(func(rayOptions *rayConfigOptions) {
		rayOptions.NumCpus = NumCpus
	})
}

func CreateRayConfig(
	redisAddress string, redisPort int, objectManagerPort int, nodeManagerPort int, gcsServerPort int,
	nodeIpAddress string, rayletIpAddress string, rayClientServerPort int, redisPassword string,
	localMode bool, opts ...RayConfigOptions) RayConfig {
	conf := RayConfig{
		RedisAddress:        redisAddress,
		RedisPort:           redisPort,
		ObjectManagerPort:   objectManagerPort,
		NodeManagerPort:     nodeManagerPort,
		GcsServerPort:       gcsServerPort,
		NodeIpAddress:       nodeIpAddress,
		RayletIpAddress:     rayletIpAddress,
		RayClientServerPort: rayClientServerPort,
		RedisPassword:       redisPassword,
		RayOpts:             DefaultRayConfigOptions,
		LocalMode:           localMode,
	}
	// Setting options
	for _, op := range opts {
		op.apply(&conf.RayOpts)
	}

	return conf
}

func DefaultRayConfig() RayConfig {
	conf := RayConfig{
		RedisAddress:        "127.0.0.1",
		RedisPort:           6379,
		ObjectManagerPort:   12345,
		NodeManagerPort:     12346,
		GcsServerPort:       12341,
		NodeIpAddress:       "0.0.0.0",
		RayletIpAddress:     "127.0.0.1",
		RayClientServerPort: 12342,
		RedisPassword:       "test",
		RayOpts:             DefaultRayConfigOptions,
	}
	return conf
}
