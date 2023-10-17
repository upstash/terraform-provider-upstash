package topic

import "github.com/upstash/terraform-provider-upstash/upstash/client"

func getTopic(c *client.UpstashClient, topicName string) (topic QStashTopic, err error) {

	resp, err := c.SendGetRequest(c.GetQstashEndpointV2()+"/topics/"+topicName, "Get QStash Topic", true)

	if err != nil {
		return topic, err
	}

	err = resp.ToJSON(&topic)
	return topic, err
}

func createTopic(c *client.UpstashClient, topicName string, body UpdateQStashTopicEndpoints) (err error) {
	_, err = c.SendPostRequest(c.GetQstashEndpointV2()+"/topics/"+topicName+"/endpoints", body, "Create QStash Topic", true)
	return err
}

func addEndpointsToTopic(c *client.UpstashClient, topicName string, body UpdateQStashTopicEndpoints) (err error) {
	_, err = c.SendPostRequest(c.GetQstashEndpointV2()+"/topics/"+topicName+"/endpoints", body, "Add QStash Endpoint for Topic", true)
	return err
}

func deleteEndpointsFromTopic(c *client.UpstashClient, topicName string, body UpdateQStashTopicEndpoints) (err error) {
	return c.SendDeleteRequest(c.GetQstashEndpointV2()+"/topics/"+topicName+"/endpoints", body, "Delete QStash Endpoints", true)
}

func deleteTopic(c *client.UpstashClient, topicName string) (err error) {
	return c.SendDeleteRequest(c.GetQstashEndpointV2()+"/topics/"+topicName, nil, "Delete QStash Topic", true)
}
