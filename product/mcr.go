package product

import (
	"fmt"

	"github.com/Mirantis/launchpaddependencies/implementation"
)

const (
	MCRComponentName = "MCR"
)

// MCR interface
type MCR struct {
	Version string
}

// Name for the Component
func (mcr MCR) Name() string {
	return MCRComponentName
}

// Provides labels of dependencies provided by MCR
func (mcr MCR) Provides() []string {
	return []string{
		fmt.Sprintf("%s:%s", RequiresKeyProduct, MCRComponentName),
		fmt.Sprintf("%s:%s", implementation.RequiresKeyImplementation, implementation.DockerComponentName),
		fmt.Sprintf("%s:%s", implementation.RequiresKeyImplementation, "docker-ee"),
	}
}
