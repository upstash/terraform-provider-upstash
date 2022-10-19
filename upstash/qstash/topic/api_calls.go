package topic

import (
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func createTopic(c *client.UpstashClient, body createQstashTopicRequest) (topic QstashTopic, err error) {
	resp, err := c.SendPostRequest(c.GetQstashEndpoint()+"/topics", body, "Create QStash Topic")

	if err != nil {
		return topic, err
	}

	err = resp.ToJSON(&topic)
	return topic, err
}

func deleteTopic(c *client.UpstashClient, topicId string) (err error) {
	return c.SendDeleteRequest(c.GetQstashEndpoint()+"/topics/"+topicId, nil, "Delete QStash Topic")
}

func getTopic(c *client.UpstashClient, topicId string) (topic QstashTopic, err error) {

	resp, err := c.SendGetRequest(c.GetQstashEndpoint()+"/topics/"+topicId, "Get QStash Topic")

	if err != nil {
		return topic, err
	}

	err = resp.ToJSON(&topic)
	return topic, err
}

func updateTopic(c *client.UpstashClient, topicId string, body UpdateQstashTopic) (err error) {
	_, err = c.SendPutRequest(c.GetQstashEndpoint()+"/topics/"+topicId, body, "Update QStash Topic")
	return err
}
