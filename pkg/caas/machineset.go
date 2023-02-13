package caas

import (
	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const machineSetKind = "MachineSet"

func createMachineSet(cluster mcaasapi.Cluster, namespace string) []clusterv1beta1.MachineSet {
	msList := []clusterv1beta1.MachineSet{}

	for _, msd := range cluster.MachineSetsDetail {

		isControlPlane := "false"
		if contains(msd.MachineRoles, "controlplane") {
			isControlPlane = "true"
		}

		msLabels := map[string]string{
			"cluster.x-k8s.io/cluster-name":  cluster.Name,
			"cluster.x-k8s.io/control-plane": isControlPlane,
			"machineSet":                     msd.Name,
			"nodeType":                       msd.Name,
		}

		ms := clusterv1beta1.MachineSet{
			TypeMeta: v1.TypeMeta{
				Kind:       machineSetKind,
				APIVersion: capiApiVersion,
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      "",
				Namespace: namespace,
				Labels:    msLabels,
			},
		}

		msList = append(msList, ms)
	}

	return msList
}
