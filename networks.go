package goCompose

type Network struct {
	Name string
	Config *NetworkConfig
}

type NetworkDriver string

const Bridge NetworkDriver = "bridge"

type NetworkConfig struct {
	External bool
	Driver NetworkDriver
}
