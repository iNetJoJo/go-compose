package goCompose

type Volume struct {
	Name string `json:"name"`
}


//TODO check if volume exists in compose
type ServiceVolume struct {
	Volume *Volume `json:"volume"`
	Path string `json:"path"`
}
