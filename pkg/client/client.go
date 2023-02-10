package client

import (
	"context"

	"github.com/HewlettPackard/hpegl-containers-go-sdk/pkg/mcaasapi"
)

func GetClusterDetails(id string, spaceID string) (mcaasapi.Cluster, error) {
	var c mcaasapi.APIClient
	field := "spaceID eq " + spaceID
	cluster, _, err := c.ClustersApi.V1ClustersIdGet(context.Background(), id, field, nil)
	if err != nil {
		return mcaasapi.Cluster{}, err
	}

	return cluster, nil

}
