package implementation

const (
	DockerComponentName = "docker"
)

// Docker host implementation
type Docker struct {

}

// Name for the Component
func (d Docker) Name() string {
	return DockerComponentName
}
