package cluster

import (
	"context"
	"fmt"
)

// Cluster collection of Hosts and Products
type Cluster struct {
	Hosts    Hosts
	Products Components
}

// ---- helper dependency functions

// ClusterDepends_ComponentProvides returns the name of component that provides a requirement (string)
func ClusterDepends_ComponentProvides(ctx context.Context, c Cluster, requires string) (string, error) {
	for _, c := range c.Products {
		for _, provides := range c.Provides() {
			if provides == requires {
				return c.Name(), nil
			}
		}
	}

	return "", fmt.Errorf("Cluster Dependency Fail: No component provides '%s'", requires)
}

// ClusterDepends_HostRoleMin Confirms that at least a certain number of hosts with a role are in the cluster
func ClusterDepends_HostRoleMin(ctx context.Context, c Cluster, role string, min int) ([]string, error) {
	hs := []string{}

	for _, h := range c.Hosts {
		if h.Role == role {
			hs = append(hs, h.Name())
		}
	}

	if len(hs) < min {
		return hs, fmt.Errorf("Cluster Dependency Fail: Not enough hosts with role '%s' [%d]", role, min)
	}
	return hs, nil
}
