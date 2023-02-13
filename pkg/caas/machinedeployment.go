package caas

import (
	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const (
	machineDeploymentKind       = "MachineDeployment"
	machineDeploymentApiVersion = "cluster.x-k8s.io/v1beta1"
	excludeNodeDraining         = "true"
)

func createMachineDeployment(cluster mcaasapi.Cluster, namespace string) []clusterv1beta1.MachineDeployment {
	mdList := []clusterv1beta1.MachineDeployment{}
	bootstrapDataSecretName := "bootstrap"

	for _, msd := range cluster.MachineSetsDetail {

		isControlPlane := "false"
		if contains(msd.MachineRoles, "controlplane") {
			isControlPlane = "true"
		}

		mdLabels := map[string]string{
			"cluster.x-k8s.io/cluster-name":  cluster.Name,
			"cluster.x-k8s.io/control-plane": isControlPlane,
			"machineSet":                     msd.Name,
			"nodeType":                       msd.Name,
		}

		md := clusterv1beta1.MachineDeployment{
			TypeMeta: v1.TypeMeta{
				Kind:       machineDeploymentKind,
				APIVersion: machineDeploymentApiVersion,
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      cluster.Name + "-" + msd.Name,
				Namespace: namespace,
				Labels:    mdLabels,
			},
			Spec: clusterv1beta1.MachineDeploymentSpec{
				ClusterName: cluster.Name,
				Template: clusterv1beta1.MachineTemplateSpec{
					ObjectMeta: clusterv1beta1.ObjectMeta{
						Annotations: map[string]string{
							"machine.cluster.x-k8s.io/exclude-node-draining": excludeNodeDraining,
						},
						Labels: mdLabels,
					},
					Spec: clusterv1beta1.MachineSpec{
						ClusterName: cluster.Name,
						Bootstrap: clusterv1beta1.Bootstrap{
							DataSecretName: &bootstrapDataSecretName,
						},
						InfrastructureRef: corev1.ObjectReference{
							APIVersion: kubeMachineTemplateApiVersion,
							Kind:       kubeMachineTemplateKind,
							Name:       getKmtName(msd.Machines[0].Name),
						},
					},
				},
			},
		}
		mdList = append(mdList, md)
	}
	return mdList
}
