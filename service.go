package goCompose

type Service struct {
	Build         string `json:"build,omitempty"`
	Image         string `json:"image,omitempty"`
	Name          string `json:"name,omitempty"`
	ContainerName string `json:"container_name,omitempty"`
	CGroupParent  string `json:"c_group_parent,omitempty"`
	Command       []string `json:"command,omitempty"`
	DeployConfig  *ServiceDeployConfig `json:"deploy_config,omitempty"`
	Configs       []string `json:"configs,omitempty"`
	DependsOn     []*Service `json:"depends_on,omitempty"`
	Ports         []*Port                `json:"ports,omitempty"`
	Volumes       []*ServiceVolume       `json:"volumes,omitempty"`
	Networks      []*Network `json:"networks,omitempty"`
	Environment   EnvironmentVariable `json:"environment,omitempty"`
}

type EnvironmentVariable struct {
	Name string `json:"name"`
	Value string `json:"value"`
}
