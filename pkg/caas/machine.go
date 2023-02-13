package caas

import (
	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	clusterv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const machineKind = "Machine"

func createMachine(cluster mcaasapi.Cluster, namespace string) []clusterv1beta1.Machine {
	mList := []clusterv1beta1.Machine{}
	//TODO: add 2 for loops
	return mList
}
