package goCompose

type ServiceDeployConfig struct {
	Mode deployMode
	Labels map[string]string
	Replicas int
	Placement deployConfigPlacement
	deployResources
}

type endPointMode string

const EndpointVIP endPointMode = "vip"
const EndpointDnsrr endPointMode = "dnsrr"


type deployMode string

const ModeGlobal deployMode = "global"
const ModeReplicated deployMode = "replicated"

type deployConfigPlacement struct {
	MaxReplicasPerNode int
	Constraints []string
	preferences map[string]string
}

type deployResources struct {
	Limits resources
	Reservations resources
}

type resources struct {
	Cpus float32
	Memory string
}


type restartPolicy string

const OnFailure restartPolicy = "on-failure"
const None restartPolicy = "none"
const Any restartPolicy = "any"

type deployRestartPolicy struct {
	Condition restartPolicy
	Delay int
	MaxAttempts int
	Window int
}


