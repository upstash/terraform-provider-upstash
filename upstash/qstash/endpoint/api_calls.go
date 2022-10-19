package endpoint

import "github.com/upstash/terraform-provider-upstash/upstash/client"

func getEndpoint(c *client.UpstashClient, endpointId string) (endpoint QstashEndpoint, err error) {

	resp, err := c.SendGetRequest(c.GetQstashEndpoint()+"/endpoints/"+endpointId, "Get QStash Endpoint")

	if err != nil {
		return endpoint, err
	}

	err = resp.ToJSON(&endpoint)
	return endpoint, err
}

func createEndpoint(c *client.UpstashClient, body createQstashEndpointRequest) (endpoint QstashEndpoint, err error) {
	resp, err := c.SendPostRequest(c.GetQstashEndpoint()+"/endpoints", body, "Create QStash Endpoint")

	if err != nil {
		return endpoint, err
	}

	err = resp.ToJSON(&endpoint)
	return endpoint, err
}

func deleteEndpoint(c *client.UpstashClient, endpointId string) (err error) {
	return c.SendDeleteRequest(c.GetQstashEndpoint()+"/endpoints/"+endpointId, nil, "Delete QStash Endpoint")
}

func updateEndpoint(c *client.UpstashClient, endpointId string, body UpdateQstashEndpoint) (err error) {
	_, err = c.SendPutRequest(c.GetQstashEndpoint()+"/endpoints/"+endpointId, body, "Update Qstash Endpoint")
	return err
}
