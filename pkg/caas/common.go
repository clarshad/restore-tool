package caas

import (
	"strconv"
	"strings"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
)

const (
	capiApiVersion   = "cluster.x-k8s.io/v1beta1"
	glCaasApiVersion = "infrastructure.cluster.x-k8s.io/v1alpha2"
)

func getKmtName(n string) string {
	s := strings.Split(n, "-")
	s = s[:len(s)-1]
	return strings.Join(s, "-")
}

// SplitUrl splits url and returns protocal, endpoint and port number
func SplitUrl(url string) (string, string, int) {
	s := strings.Split(url, ":")
	p, _ := strconv.Atoi(s[2])
	return s[0], strings.Trim(s[1], "/"), p
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
