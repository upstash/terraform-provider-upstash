package topic

type Topic struct {
	TopicId          string `json:"topic_id"`
	TopicName        string `json:"name"`
	Partitions       int64  `json:"partitions"`
	MaxRetentionTime int64  `json:"max_retention_time"`
	MaxRetentionSize int64  `json:"max_retention_size"`
	MaxMessageSize   int64  `json:"max_message_size"`
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

	CreationTime int64 `json:"creation_time"`
}

type CreateTopicRequest struct {
	TopicName      string `json:"name"`
	Partitions     int64  `json:"partitions"`
	RetentionTime  int64  `json:"retention_time"`
	RetentionSize  int64  `json:"retention_size"`
	MaxMessageSize int64  `json:"max_message_size"`
	CleanupPolicy  string `json:"cleanup_policy"`
	ClusterId      string `json:"cluster_id"`
}

type ReconfigureTopic struct {
	RetentionTime  int64 `json:"retention_time"`
	RetentionSize  int64 `json:"retention_size"`
	MaxMessageSize int64 `json:"max_message_size"`
}
