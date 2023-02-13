package caas

import (
	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
)

const (
	machineDeploymentKind       = "MachineDeployment"
	machineDeploymentApiVersion = "cluster.x-k8s.io/v1beta1"
)

type MachineDeployment struct {
}

func createMachineDeployment(cluster mcaasapi.Cluster, namespace string) MachineDeployment {
	return MachineDeployment{}
}
