package caas

import (
	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const clusterKind = "Cluster"

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec `json:"spec,omitempty"`
}

// ClusterSpec ...
type ClusterSpec struct {
	ControlPlaneEndpoint APIEndpoint             `json:"controlPlaneEndpoint,omitempty"`
	InfrastructureRef    *corev1.ObjectReference `json:"infrastructureRef,omitempty"`
}

func CreateCluster(cluster mcaasapi.Cluster, namespace string) Cluster {
	_, endpoint, port := SplitUrl(cluster.ApiEndpoint)
	objRef := corev1.ObjectReference{
		Kind:       kubeClusterKind,
		Namespace:  namespace,
		Name:       cluster.Name,
		APIVersion: glCaasApiVersion,
	}

	return Cluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       clusterKind,
			APIVersion: capiApiVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.Name,
			Namespace: namespace,
		},
		Spec: ClusterSpec{
			ControlPlaneEndpoint: APIEndpoint{
				Host: endpoint,
				Port: port,
			},
			InfrastructureRef: &objRef,
		},
	}
}
