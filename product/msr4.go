package product

import (
	"context"

	"github.com/Mirantis/launchpaddependencies/cluster"
)

// MSR4 interface
type MSR4 struct {
	Version string
}

// Name for the Component
func (_ MSR4) Name() string {
	return "MSR4"
}

// Provides labels of dependencies provided by MSR4
func (_ MSR4) Provides() []string {
	return []string{
		"product:MSR4",
		"implementation:docker-registry",
	}
}

// Validate the cluster Definition
func (_ MSR4) Validate(ctx context.Context, c cluster.Cluster) ([]string, error) {
	deps := []string{}

	if md, pm_err := cluster.ClusterDepends_ComponentProvides(ctx, c, "implementation:kubernetes"); pm_err != nil {
		return deps, pm_err
	} else {
		deps = append(deps, md)
	}

	return deps, nil
}
