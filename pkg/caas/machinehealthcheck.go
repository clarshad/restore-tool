package caas

import (
	"time"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	clusterv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const machineHealthCheckKind = "MachineHealthCheck"

func createMachineHealthCheck(cluster mcaasapi.Cluster, namespace string) clusterv1beta1.MachineHealthCheck {

	maxUnhealthy := intstr.IntOrString{StrVal: "99%"}
	nodeStartuptime, _ := time.ParseDuration("0s")
	timeoutUnhealthy, _ := time.ParseDuration("10m0s")

	return clusterv1beta1.MachineHealthCheck{
		TypeMeta: v1.TypeMeta{
			Kind:       machineHealthCheckKind,
			APIVersion: capiApiVersion,
		},
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name:      cluster.Name + "-mhc",
			Labels: map[string]string{
				"cluster.x-k8s.io/cluster-name": cluster.Name,
			},
		},
		Spec: clusterv1beta1.MachineHealthCheckSpec{
			ClusterName:        cluster.Name,
			MaxUnhealthy:       &maxUnhealthy,
			NodeStartupTimeout: &v1.Duration{Duration: nodeStartuptime},
			Selector: v1.LabelSelector{
				MatchLabels: map[string]string{
					"nodeType": "worker",
				},
			},
			UnhealthyConditions: []clusterv1beta1.UnhealthyCondition{
				{
					Type:    corev1.NodeConditionType("Ready"),
					Timeout: v1.Duration{Duration: timeoutUnhealthy},
					Status:  corev1.ConditionStatus("Unknown"),
				},
				{
					Type:    corev1.NodeConditionType("Ready"),
					Timeout: v1.Duration{Duration: timeoutUnhealthy},
					Status:  corev1.ConditionStatus("False"),
				},
			},
		},
	}
}
