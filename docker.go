package main

const (
	DockerComponentName = "Docker"
)

// Docker interface
type Docker struct {

}

// Name for the Component
func (d Docker) Name() string {
	return DockerComponentName
}
