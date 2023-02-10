package caas

import (
	"strings"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kubeMachineKind       = "KubeMachine"
	kubeMachineApiVersion = "infrastructure.cluster.x-k8s.io/v1alpha2"
)

// KubeMachine is the Schema for the kubemachines API
type KubeMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeMachineSpec `json:"spec,omitempty"`
}

// KubeMachineSpec ...
type KubeMachineSpec struct {
	ProviderID          *string  `json:"providerID,omitempty"`
	ProviderRef         string   `json:"providerRef"`
	OsImage             string   `json:"osImage"`
	OsVersion           string   `json:"osVersion"`
	ComputeInstanceType string   `json:"computeInstanceType,omitempty"`
	StorageInstanceType string   `json:"storageInstanceType,omitempty"`
	Size                string   `json:"size"`
	Networks            []string `json:"networks"`
	Roles               []string `json:"roles"`
}

func CreateKubeMachine(cluster mcaasapi.Cluster, namespace string, dmn map[string]string) []KubeMachine {
	kmlist := []KubeMachine{}

	for _, msd := range cluster.MachineSetsDetail {
		var machine_networks []string
		var isControlPlane string

		if contains(msd.MachineRoles, "controlplane") {
			machine_networks = strings.Split(dmn["controlplane"], ", ")
			isControlPlane = "true"
		} else {
			machine_networks = strings.Split(dmn["worker"], ", ")
			isControlPlane = "false"
		}

		for _, m := range msd.Machines {
			km := KubeMachine{
				TypeMeta: metav1.TypeMeta{
					Kind:       kubeMachineKind,
					APIVersion: kubeMachineApiVersion,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      m.Name,
					Namespace: namespace,
					Labels: map[string]string{
						"cluster":                        cluster.Name,
						"cluster.x-k8s.io/cluster-name":  cluster.Name,
						"cluster.x-k8s.io/control-plane": isControlPlane,
						"platformClusterID":              cluster.Id,
						"machineSet":                     msd.Name,
						"nodeType":                       string(msd.MachineRoles[0]),
					},
				},
				Spec: KubeMachineSpec{
					ComputeInstanceType: msd.ComputeInstanceType,
					OsImage:             msd.OsImage,
					OsVersion:           msd.OsVersion,
					ProviderRef:         msd.MachineProvider,
					Size:                msd.Size,
					StorageInstanceType: msd.StorageInstanceType,
					Networks:            machine_networks,
					Roles:               ConvertToStringSlice(msd.MachineRoles),
				},
			}

			kmlist = append(kmlist, km)
		}
	}
	return kmlist
}

func contains(s []mcaasapi.MachineRolesType, str string) bool {
	for _, v := range s {
		if string(v) == str {
			return true
		}
	}
	return false
}

func ConvertToStringSlice(s []mcaasapi.MachineRolesType) []string {
	list := []string{}
	for _, v := range s {
		list = append(list, string(v))
	}
	return list
}
