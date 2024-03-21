package main

import "context"

// Phase an action step in a launchpad operation
type Phase interface {
	// Label for logging the phase (not called name, just to avoid confusing interfaces)
	Label() string
	// DependsOn list dependencies
	Requires() []string
	// Run the phase
	Run(context.Context) error
}
