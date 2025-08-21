package topic

import "github.com/upstash/terraform-provider-upstash/v2/upstash/qstash/endpoint"

type createQstashTopicRequest struct {
	Name string `json:"name"`
}

type QstashTopic struct {
	Name      string                    `json:"name"`
	TopicId   string                    `json:"topicId"`
	Endpoints []endpoint.QstashEndpoint `json:"endpoints"`
}

type UpdateQstashTopic struct {
	Name string `json:"name"`
}
