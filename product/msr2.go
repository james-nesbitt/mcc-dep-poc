package product

import (
	"fmt"

	"github.com/Mirantis/launchpaddependencies/implementation"
)

const (
	MSR2ComponentName = "MSR2"
)

// MKE interface
type MSR2 struct {
	Version string
}

// Name for the Component
func (_ MSR2) Name() string {
	return MSR2ComponentName
}

// Provides labels of dependencies provided by MSR2
func (_ MSR2) Provides() []string {
	return []string{
		fmt.Sprintf("%s:%s", RequiresKeyProduct ,MSR2ComponentName),
		fmt.Sprintf("%s:%s", implementation.RequiresKeyImplementation, implementation.KubernetesComponentName ),
	}
}
