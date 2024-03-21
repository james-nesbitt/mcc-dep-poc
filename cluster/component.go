package cluster

import "context"

/**
* Component
*
* A part of a cluster which can have dependencies, or be a dependency
*
* Products:
*  Products are components which interact with the cluster, installing
*  onto the hosts of the cluster, and providing interfaces
*
* Interface:
*  Interfaces are components of the cluster which provide functionality
*  which may be needed by other components. Interfaces exists to allow
*  different other components such as products to provide the same
*  functionality.
 */

// Components collection
type Components []Component

// Component a cluster component
type Component interface {
	// Name of the Component, which is used for auditing, but perhaps
	//   also as the basis for a dependency.
	Name() string
	// Provides a list of string labels which this Component provides,
	//   which can be used for auditing, but is focused on simplifying
	//   some early dependency testing.
	Provides() []string
	// Validate that the cluster definition meets the needs of the
	//   product, and return a list of dependencies
	Validate(context.Context, Cluster) ([]string, error)
}
