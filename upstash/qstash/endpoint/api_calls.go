package endpoint

import "github.com/upstash/terraform-provider-upstash/upstash/client"

var QSTASH_API_ENDPOINT = "https://qstash.upstash.io/v1"

func getEndpoint(c *client.UpstashClient, endpointId string) (endpoint QstashEndpoint, err error) {

	resp, err := c.SendGetRequest(QSTASH_API_ENDPOINT+"/endpoints/"+endpointId, "Get QStash Endpoint")

	if err != nil {
		return endpoint, err
	}

	err = resp.ToJSON(&endpoint)
	return endpoint, err
}

func createEndpoint(c *client.UpstashClient, body createQstashEndpointRequest) (endpoint QstashEndpoint, err error) {
	resp, err := c.SendPostRequest(QSTASH_API_ENDPOINT+"/endpoints", body, "Create QStash Endpoint")

	if err != nil {
		return endpoint, err
	}

	err = resp.ToJSON(&endpoint)
	return endpoint, err
}

func deleteEndpoint(c *client.UpstashClient, endpointId string) (err error) {
	return c.SendDeleteRequest(QSTASH_API_ENDPOINT+"/endpoints/"+endpointId, nil, "Delete QStash Endpoint")
}
