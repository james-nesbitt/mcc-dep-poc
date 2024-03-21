package product

import (
	"fmt"

	"github.com/Mirantis/launchpaddependencies/implementation"
)

const (
	MKE3ComponentName = "MKE3"
)

// MKE interface
type MKE3 struct {
	Version string
}

// Name for the Component
func (_ MKE3) Name() string {
	return MKE3ComponentName
}

// Provides labels of dependencies provided by MKE3
func (_ MKE3) Provides() []string {
	return []string{
		fmt.Sprintf("%s:%s", RequiresKeyProduct ,MKE3ComponentName),
		fmt.Sprintf("%s:%s", implementation.RequiresKeyImplementation, implementation.KubernetesComponentName ),
	}
}
