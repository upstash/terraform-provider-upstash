package cluster

import (
	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func createCluster(c *client.UpstashClient, body CreateClusterRequest) (cluster Cluster, err error) {

	resp, err := c.SendPostRequest("/v2/kafka/cluster", body, "Create Kafka Cluster", false)

	if err != nil {
		return cluster, err
	}

	err = resp.ToJSON(&cluster)
	return cluster, err
}

func getCluster(c *client.UpstashClient, clusterId string) (cluster Cluster, err error) {

	resp, err := c.SendGetRequest("/v2/kafka/cluster/"+clusterId, "Get Kafka Cluster", false)

	if err != nil {
		return cluster, err
	}

	err = resp.ToJSON(&cluster)
	return cluster, err

}

func renameCluster(c *client.UpstashClient, clusterId string, newName string) (err error) {

	_, err = c.SendPostRequest("/v2/kafka/rename-cluster/"+clusterId, req.Param{"name": newName}, "Rename Kafka Cluster", false)

	return err

}

func deleteCluster(c *client.UpstashClient, clusterId string) (err error) {

	return c.SendDeleteRequest("/v2/kafka/cluster/"+clusterId, nil, "Delete Kafka Cluster", false)
}
