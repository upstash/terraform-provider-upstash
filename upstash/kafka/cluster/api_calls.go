package cluster

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

const api_endpoint = client.UPSTASH_API_ENDPOINT

func createCluster(c *client.UpstashClient, body CreateClusterRequest) (cluster Cluster, err error) {
	resp, err := req.Post(api_endpoint+"/v2/kafka/cluster",
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
		req.BodyJSON(body),
	)
	if err != nil {
		return cluster, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return cluster, errors.New("Create cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&cluster)
	return cluster, err

}

func getCluster(c *client.UpstashClient, clusterId string) (cluster Cluster, err error) {
	resp, err := req.Get(api_endpoint+"/v2/kafka/cluster/"+clusterId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return cluster, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return cluster, errors.New("Get cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&cluster)
	return cluster, err

}

func renameCluster(c *client.UpstashClient, clusterId string, newName string) (err error) {

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": utils.BasicAuth(c.Email, c.Apikey),
	}

	param := req.Param{
		"name": newName,
	}

	resp, err := req.Post(api_endpoint+"/v2/kafka/rename-cluster/"+clusterId, req.BodyJSON(param), header)

	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Renaming cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return nil
}

func deleteCluster(c *client.UpstashClient, clusterId string) (err error) {
	resp, err := req.Delete(api_endpoint+"/v2/kafka/cluster/"+clusterId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Delete cluster failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())

	}
	return err
}
