package product

import (
	"context"

	"github.com/Mirantis/launchpaddependencies/cluster"
)

// MKE interface
type MSR2 struct {
	Version string
}

// Name for the Component
func (_ MSR2) Name() string {
	return "MSR2"
}

// Provides labels of dependencies provided by MSR2
func (_ MSR2) Provides() []string {
	return []string{
		"product:MSR2",
		"implementation:docker-registry",
	}
}

// Validate the cluster Definition
func (_ MSR2) Validate(ctx context.Context, c cluster.Cluster) ([]string, error) {
	deps := []string{}

	if md, pm_err := cluster.ClusterDepends_ComponentProvides(ctx, c, "product:mke3"); pm_err != nil {
		return deps, pm_err
	} else {
		deps = append(deps, md)
	}

	if hrs, hr_err := cluster.ClusterDepends_HostRoleMin(ctx, c, "msr", 1); hr_err.Error != nil {
		return deps, hr_err
	} else {
		for _, hr := range hrs {
			deps = append(deps, hr)
		}
	}

	return deps, nil
}
