package topic

type Topic struct {
	TopicId          string `json:"topic_id"`
	TopicName        string `json:"topic_name"`
	Partitions       int    `json:"partitions"`
	MaxRetentionTime int    `json:"retention_time"`
	MaxRetentionSize int    `json:"retention_size"`
	MaxMessageSize   int    `json:"max_message_size"`
	CleanupPolicy    string `json:"cleanup_policy"`
	ClusterId        string `json:"cluster_id"`
	Region           string `json:"region"`
	State            string `json:"state"`
	MultiZone        bool   `json:"multizone"`
	TcpEndpoint      string `json:"tcp_endpoint"`
	RestEndpoint     string `json:"rest_endpoint"`
	Username         string `json:"username"`
	EncodedUsername  string `json:"encoded_username"`
	Password         string `json:"password"`

	CreationTime int `json:"creation_time"`
}

type CreateTopicRequest struct {
	TopicName      string `json:"name"`
	Partitions     int    `json:"partitions"`
	RetentionTime  int    `json:"retention_time"`
	RetentionSize  int    `json:"retention_size"`
	MaxMessageSize int    `json:"max_message_size"`
	CleanupPolicy  string `json:"cleanup_policy"`
	ClusterId      string `json:"cluster_id"`
}

type ReconfigureTopic struct {
	RetentionTime  int `json:"retention_time"`
	RetentionSize  int `json:"retention_size"`
	MaxMessageSize int `json:"max_message_size"`
}
