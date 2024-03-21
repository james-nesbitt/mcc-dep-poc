package product

import (
	"context"

	"github.com/Mirantis/launchpaddependencies/cluster"
)

// MKE4 interface
type MKE4 struct {
	Version string
}

// Name for the Component
func (_ MKE4) Name() string {
	return "MKE4"
}

// Provides labels of dependencies provided by MKE4
func (_ MKE4) Provides() []string {
	return []string{
		"product:MKE4",
	}
}

// Validate the cluster Definition
func (_ MKE4) Validate(ctx context.Context, c cluster.Cluster) ([]string, error) {
	deps := []string{}

	if md, pm_err := cluster.ClusterDepends_ComponentProvides(ctx, c, "implementation:kubernetes"); pm_err != nil {
		return deps, pm_err
	} else {
		deps = append(deps, md)
	}

	return deps, nil
}
