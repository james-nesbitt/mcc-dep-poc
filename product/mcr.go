package product

import (
	"context"

	"github.com/Mirantis/launchpaddependencies/cluster"
)

// MCR interface
type MCR struct {
	Version string
}

// Name for the Component
func (_ MCR) Name() string {
	return "MCR"
}

// Provides labels of dependencies provided by MCR
func (_ MCR) Provides() []string {
	return []string{
		"product:MCR",
		"implementation:docker-swarm",
		"implementation:docker-host",
	}
}

// Validate the cluster Definition
func (_ MCR) Validate(_ context.Context, c cluster.Cluster) ([]string, error) {
	return []string{}, nil
}
