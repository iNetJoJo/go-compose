package goCompose

type Service struct {
	Build         string //done
	Image         string //done
	Name          string //done
	ContainerName string //done
	CGroupParent  string //done
	Command       []string //done
	DeployConfig  *ServiceDeployConfig //todo
	Configs       []string //todo
	DependsOn     []*Service // done
	Ports         []*Port // done
	Volumes       []*ServiceVolume //done
	Networks      []*Network //done
	Environment   map[string]interface{} //done
}
