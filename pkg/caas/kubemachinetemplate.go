package caas

import (
	"strings"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kubeMachineTemplateKind       = "KubeMachine"
	KubeMachineTemplateApiVersion = "infrastructure.cluster.x-k8s.io/v1alpha2"
)

// KubeMachineTemplate is the Schema for the kubemachinetemplates
type KubeMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeMachineTemplateSpec `json:"spec,omitempty"`
}

type KubeMachineTemplateSpec struct {
	Template KubeMachineTemplateResource `json:"template"`
}

type KubeMachineTemplateResource struct {
	Spec KubeMachineSpec `json:"spec"`
}

func CreateKubeMachineTemplate(cluster mcaasapi.Cluster, namespace string) []KubeMachineTemplate {
	kmtlist := []KubeMachineTemplate{}

	for _, msd := range cluster.MachineSetsDetail {
		name := getKmtName(msd.Machines[0].Name)

		kmt := KubeMachineTemplate{
			TypeMeta: metav1.TypeMeta{
				Kind:       kubeMachineTemplateKind,
				APIVersion: KubeMachineTemplateApiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: KubeMachineTemplateSpec{
				KubeMachineTemplateResource{
					KubeMachineSpec{
						ComputeInstanceType: msd.ComputeInstanceType,
						OsImage:             msd.OsImage,
						OsVersion:           msd.OsVersion,
						ProviderRef:         msd.MachineProvider,
						Size:                msd.Size,
						StorageInstanceType: msd.StorageInstanceType,
						Networks:            []string{},
						Roles:               ConvertToStringSlice(msd.MachineRoles),
					},
				},
			},
		}

		kmtlist = append(kmtlist, kmt)
	}

	return kmtlist
}

func getKmtName(n string) string {
	s := strings.Split(n, "-")
	s = s[:len(s)-1]
	return strings.Join(s, "-")
}
