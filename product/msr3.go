package product

import (
	"context"

	"github.com/Mirantis/launchpaddependencies/cluster"
)

// MSR3 interface
type MSR3 struct {
	Version string
}

// Name for the Component
func (_ MSR3) Name() string {
	return "MSR3"
}

// Provides labels of dependencies provided by MSR3
func (_ MSR3) Provides() []string {
	return []string{
		"product:MSR3",
		"implementation:docker-registry",
	}
}

// Validate the cluster Definition
func (_ MSR3) Validate(ctx context.Context, c cluster.Cluster) ([]string, error) {
	deps := []string{}

	if md, pm_err := cluster.ClusterDepends_ComponentProvides(ctx, c, "implementation:kubernetes"); pm_err != nil {
		return deps, pm_err
	} else {
		deps = append(deps, md)
	}

	return deps, nil
}
