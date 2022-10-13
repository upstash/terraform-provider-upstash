package endpoint

type createQstashEndpointRequest struct {
	Url       string `json:"url"`
	TopicName string `json:"topicName,omitempty"`
	TopicId   string `json:"topicId,omitempty"`
}

type QstashEndpoint struct {
	Url        string `json:"url"`
	TopicId    string `json:"topicId"`
	EndpointId string `json:"endpointId"`
}
