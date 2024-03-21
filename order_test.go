package main_test

import (
	"testing"

	launchpad "github.com/Mirantis/launchpaddependencies"
	"github.com/Mirantis/launchpaddependencies/product"
)

func Cluster_MCRMKEMSR() launchpad.Cluster {
	return launchpad.Cluster{
		Hosts: launchpad.Hosts{
			launchpad.Host{
				Role: "manager",
				OS:   "Ubuntu",
			},
			launchpad.Host{
				Role: "worker",
				OS:   "Ubuntu",
			},
			launchpad.Host{
				Role: "worker",
				OS:   "Ubuntu",
			},
		},
		Products: launchpad.Components{
			product.MCR{
				Version: "20.10.15",
			},
			product.MKE3{
				Version: "3.7.5",
			},
			product.MSR2{
				Version: "2.9.16",
			},
		},
	}
}

func TestDependencyOrder(t *testing.T) {
}
