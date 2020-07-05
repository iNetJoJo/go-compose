package goCompose

type Volume struct {
	Name string
}


//TODO check if volume exists in compose
type ServiceVolume struct {
	Volume *Volume
	Path string
}
