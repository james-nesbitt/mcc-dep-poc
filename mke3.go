package main

const (
	MKEComponentName = "MKE"
)

// MKE interface
type MKE struct {

}

// Name for the Component
func (d MKE) Name() string {
	return MKEComponentName
}
