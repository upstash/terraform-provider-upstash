package topic

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

const api_endpoint = client.UPSTASH_API_ENDPOINT

func createTopic(c *client.UpstashClient, body CreateTopicRequest) (topic Topic, err error) {
	resp, err := req.Post(api_endpoint+"/v2/kafka/topic",
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
		req.BodyJSON(body),
	)
	if err != nil {
		return topic, err
	}

	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return topic, errors.New("Create kafka topic failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&topic)
	return topic, err
}

func getTopic(c *client.UpstashClient, topicId string) (topic Topic, err error) {
	resp, err := req.Get(api_endpoint+"/v2/kafka/topic/"+topicId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
	)

	if err != nil {
		return topic, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return topic, errors.New("Get kafka topic failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&topic)
	return topic, err
}

func reconfigureKafkaTopic(c *client.UpstashClient, topicId string, body ReconfigureTopic) (err error) {
	resp, err := req.Post(api_endpoint+"/v2/kafka/update-topic/"+topicId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
		req.BodyJSON(body),
	)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Reconfiguring kafka topic failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	return nil
}

func deleteTopic(c *client.UpstashClient, topicId string) (err error) {
	resp, err := req.Delete(api_endpoint+"/v2/kafka/topic/"+topicId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Delete kafka topic failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())

	}
	return err
}
