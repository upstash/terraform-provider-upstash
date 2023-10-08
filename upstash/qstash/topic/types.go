package topic

type QStashEndpoint struct {
	Url string `json:"url"`
}

type QStashTopic struct {
	Name      string           `json:"name"`
	CreatedAt int64            `json:"created_at"`
	UpdatedAt int64            `json:"updated_at"`
	Endpoints []QStashEndpoint `json:"endpoints"`
}

type UpdateQStashTopicEndpoints struct {
	Endpoints []QStashEndpoint `json:"endpoints"`
}
