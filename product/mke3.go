package product

import (
	"context"

	"github.com/Mirantis/launchpaddependencies/cluster"
)

// MKE interface
type MKE3 struct {
	Version string
}

// Name for the Component
func (_ MKE3) Name() string {
	return "MKE3"
}

// Provides labels of dependencies provided by MKE3
func (_ MKE3) Provides() []string {
	return []string{
		"product:MKE3",
		"implementation:kubernetes",
	}
}

// Validate the cluster Definition
func (_ MKE3) Validate(ctx context.Context, c cluster.Cluster) ([]string, error) {
	deps := []string{}

	if md, pm_err := cluster.ClusterDepends_ComponentProvides(ctx, c, "product:mcr"); pm_err != nil {
		return deps, pm_err
	} else {
		deps = append(deps, md)
	}

	if hrs, hr_err := cluster.ClusterDepends_HostRoleMin(ctx, c, "manager", 1); hr_err.Error != nil {
		return deps, hr_err
	} else {
		for _, hr := range hrs {
			deps = append(deps, hr)
		}
	}

	return deps, nil
}
