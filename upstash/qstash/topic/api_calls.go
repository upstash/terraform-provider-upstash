package topic

import (
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

var QSTASH_API_ENDPOINT = "https://qstash.upstash.io/v1"

func createTopic(c *client.UpstashClient, body createQstashTopicRequest) (topic QstashTopic, err error) {
	resp, err := c.SendPostRequest(QSTASH_API_ENDPOINT+"/topics", body, "Create QStash Topic")

	if err != nil {
		return topic, err
	}

	err = resp.ToJSON(&topic)
	return topic, err
}

func deleteTopic(c *client.UpstashClient, topicId string) (err error) {
	return c.SendDeleteRequest(QSTASH_API_ENDPOINT+"/topics/"+topicId, nil, "Delete QStash Topic")
}

func getTopic(c *client.UpstashClient, topicId string) (topic QstashTopic, err error) {

	resp, err := c.SendGetRequest(QSTASH_API_ENDPOINT+"/topics/"+topicId, "Get QStash Topic")

	if err != nil {
		return topic, err
	}

	err = resp.ToJSON(&topic)
	return topic, err
}
