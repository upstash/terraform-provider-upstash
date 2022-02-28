package cluster

type Cluster struct {
	ClusterId   string `json:"cluster_id"`
	ClusterName string `json:"name"`
	Region      string `json:"region"`

	Type                 string `json:"type"`
	MultiZone            bool   `json:"multizone"`
	TcpEndpoint          string `json:"tcp_endpoint"`
	RestEndpoint         string `json:"rest_endpoint"`
	State                string `json:"state"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	MaxRetentionSize     int64  `json:"max_retention_size"`
	MaxRetentionTime     int64  `json:"max_retention_time"`
	MaxMessagesPerSecond int    `json:"max_messages_per_second"`
	MaxMessageSize       int64  `json:"max_message_size"`
	MaxPartitions        int    `json:"max_partitions"`

	CreationTime int64 `json:"creation_time"`
}

type CreateClusterRequest struct {
	ClusterName string `json:"name"`
	Region      string `json:"region"`
	MultiZone   bool   `json:"multizone"`
}
