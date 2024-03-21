package cluster

import (
	"context"
	"fmt"
)

type Hosts []Host

type Host struct {
	ID string
	
	Role string
	OS string
}

// Name component label
func (h Host) Name() string {
	return fmt.Sprintf("%s:%s", "host", h.ID)
}

// Provides labels of dependencies provided by MSR2
func (_ Host) Provides() []string {
	return []string{
		"component:host",
	}
}

// Validate the cluster Definition
func (_ Host) Validate(ctx context.Context, c Cluster) ([]string, error) {
	return []string{}, nil
}
