package caas

import (
	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kubeBootstrapConfigKind       = "KubeBootstrapConfig"
	KubeBootstrapConfigApiVersion = "infrastructure.cluster.x-k8s.io/v1alpha2"
)

// KubeBootstrapConfig is the Schema for the kubebootstrapconfigs
type KubeBootstrapConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeBootstrapConfigSpec   `json:"spec,omitempty"`
	Status            KubeBootstrapConfigStatus `json:"status,omitempty"`
}

// KubeBootstrapConfigSpec ...
type KubeBootstrapConfigSpec struct {
	Cluster string `json:"cluster"`
	Machine string `json:"machine"`
}

// KubeBootstrapConfigStatus ...
type KubeBootstrapConfigStatus struct {
	Created bool `json:"created"`
	Ready   bool `json:"ready"`
}

func CreateKubeBootstrapConfig(cluster mcaasapi.Cluster, namespace string) []KubeBootstrapConfig {
	kbcList := []KubeBootstrapConfig{}

	for _, msd := range cluster.MachineSetsDetail {
		for _, m := range msd.Machines {
			kbc := KubeBootstrapConfig{
				TypeMeta: metav1.TypeMeta{
					Kind:       kubeBootstrapConfigKind,
					APIVersion: KubeBootstrapConfigApiVersion,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      cluster.Name + "-" + m.Name,
					Namespace: namespace,
				},
				Spec: KubeBootstrapConfigSpec{
					Cluster: cluster.Name,
					Machine: m.Name,
				},
				Status: KubeBootstrapConfigStatus{
					Created: true,
					Ready:   true,
				},
			}
			kbcList = append(kbcList, kbc)
		}
	}
	return kbcList
}
