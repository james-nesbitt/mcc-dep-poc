package main

const (
	K0sComponentName = "K0s"
)

// K0s interface
type K0s struct {

}

// Name for the Component
func (d K0s) Name() string {
	return K0sComponentName
}
