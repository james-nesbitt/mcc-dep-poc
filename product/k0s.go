package product

import (
	"context"

	"github.com/Mirantis/launchpaddependencies/cluster"
)

// K0s interface
type K0s struct {
	Version string
}

// Name for the Component
func (_ K0s) Name() string {
	return "K0S"
}

// Provides labels of dependencies provided by k0s
func (_ K0s) Provides() []string {
	return []string{
		"product:K0S",
		"implementation:kubernetes",
	}
}

// Validate the cluster Definition
func (_ K0s) Validate(_ context.Context, c cluster.Cluster) (err error) {

	return nil
}
