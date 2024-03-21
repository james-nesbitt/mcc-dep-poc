package implementation

const (
	KubernetesComponentName = "kubernetes"
)

// Kubernetes interface
type Kubernetes struct {

}

// Name for the Component
func (d Kubernetes) Name() string {
	return KubernetesComponentName
}
