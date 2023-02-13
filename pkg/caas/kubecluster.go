package caas

import (
	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const kubeClusterKind = "KubeCluster"

type KubeCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeClusterSpec `json:"spec,omitempty"`
}

type KubeClusterSpec struct {
	Version              string      `json:"version,omitempty"`
	ProviderRef          string      `json:"providerRef"`
	DefaultStorageClass  string      `json:"defaultStorageClass,omitempty"`
	ControlPlaneEndpoint APIEndpoint `json:"controlPlaneEndpoint,omitempty"`
}

type APIEndpoint struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func CreateKubeCluster(cluster mcaasapi.Cluster, namespace string) KubeCluster {

	_, endpoint, port := SplitUrl(cluster.ApiEndpoint)
	return KubeCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       kubeClusterKind,
			APIVersion: glCaasApiVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.Name,
			Namespace: namespace,
			Labels: map[string]string{
				"platformClusterID":             string(cluster.Id),
				"cluster.x-k8s.io/cluster-name": cluster.Name,
			},
		},
		Spec: KubeClusterSpec{
			ProviderRef:         cluster.ClusterProvider,
			Version:             cluster.KubernetesVersion,
			DefaultStorageClass: cluster.DefaultStorageClass,
			ControlPlaneEndpoint: APIEndpoint{
				Host: endpoint,
				Port: port,
			},
		},
	}
}
