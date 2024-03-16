package main

const (
	MCRComponentName = "MCR"
)

// MCR interface
type MCR struct {

}

// Name for the Component
func (d MCR) Name() string {
	return MCRComponentName
}
