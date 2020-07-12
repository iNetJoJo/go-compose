package goCompose

type ServiceDeployConfig struct {
	Mode deployMode `json:"mode,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	Replicas int `json:"replicas,omitempty"`
	Placement deployConfigPlacement `json:"placement,omitempty"`
	DeployResources deployResources `json:"deploy_resources,omitempty"` 
}

type endPointMode string

const EndpointVIP endPointMode = "vip"
const EndpointDnsrr endPointMode = "dnsrr"


type deployMode string

const ModeGlobal deployMode = "global"
const ModeReplicated deployMode = "replicated"

type deployConfigPlacement struct {
	MaxReplicasPerNode int `json:"max_replicas_per_node,omitempty"`
	Constraints []string `json:"constraints,omitempty"`
	preferences map[string]string `json:"preferences,omitempty"`
}

type deployResources struct {
	Limits resources `json:"limits,omitempty"`
	Reservations resources `json:"reservations,omitempty"`
}

type resources struct {
	Cpus float32 `json:"cpus,omitempty"`
	Memory string `json:"memory,omitempty"`
}


type restartPolicy string

const OnFailure restartPolicy = "on-failure"
const None restartPolicy = "none"
const Any restartPolicy = "any"

type deployRestartPolicy struct {
	Condition restartPolicy `json:"condition,omitempty"`
	Delay int `json:"delay,omitempty"`
	MaxAttempts int `json:"max_attempts,omitempty"`
	Window int `json:"window,omitempty"`
}


