package product

import (
	"fmt"

	"github.com/Mirantis/launchpaddependencies/implementation"
)

const (
	K0sComponentName = "K0s"
)

// K0s interface
type K0s struct {
	Version string
}

// Name for the Component
func (d K0s) Name() string {
	return K0sComponentName
}

// Provides labels of dependencies provided by k0s
func (d K0s) Provides() []string {
	return []string{
		fmt.Sprintf("%s:%s", RequiresKeyProduct, K0sComponentName),
		fmt.Sprintf("%s:%s", implementation.RequiresKeyImplementation, implementation.KubernetesComponentName),
	}
}
