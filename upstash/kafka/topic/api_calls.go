package topic

import (
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func createTopic(c *client.UpstashClient, body CreateTopicRequest) (topic Topic, err error) {

	resp, err := c.SendPostRequest("/v2/kafka/topic", body, "Create Kafka Topic", false)

	if err != nil {
		return topic, err
	}

	err = resp.ToJSON(&topic)
	return topic, err

}

func getTopic(c *client.UpstashClient, topicId string) (topic Topic, err error) {

	resp, err := c.SendGetRequest("/v2/kafka/topic/"+topicId, "Get Kafka Topic", false)

	if err != nil {
		return topic, err
	}

	err = resp.ToJSON(&topic)
	return topic, err

}

func reconfigureKafkaTopic(c *client.UpstashClient, topicId string, body ReconfigureTopic) (err error) {

	_, err = c.SendPostRequest("/v2/kafka/update-topic/"+topicId, body, "Reconfigure Kafka Cluster", false)

	return err

}

func deleteTopic(c *client.UpstashClient, topicId string) (err error) {

	return c.SendDeleteRequest("/v2/kafka/topic/"+topicId, nil, "Delete Kafka Topic", false)

}
