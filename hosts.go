package main

type Hosts []Host

type Host struct {
	Role string
	OS string
}

func (h Host) IsWindows() bool {
	return h.OS == "windows"
}

func (h Host) IsControlPlane() bool {
	return h.Role == "manager" || h.Role == "controller"
}
