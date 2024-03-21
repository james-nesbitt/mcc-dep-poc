package main_test

import (
	"testing"

	"github.com/Mirantis/launchpaddependencies/cluster"
	"github.com/Mirantis/launchpaddependencies/product"
)

func Cluster_MCRMKEMSR() cluster.Cluster {
	return cluster.Cluster{
		Hosts: cluster.Hosts{
			cluster.Host{
				ID: "boss",
				Role: "manager",
				OS:   "Ubuntu",
			},
			cluster.Host{
				ID: "worker_1",
				Role: "worker",
				OS:   "Ubuntu",
			},
			cluster.Host{
				ID: "worker_2",
				Role: "worker",
				OS:   "Ubuntu",
			},
		},
		Products: cluster.Components{
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

// confirm that all cluster component interfaces are good
func Test_Clean(t *testing.T) {

	Cluster_MCRMKEMSR()
	
}
