package upstash

type Database struct {
	DatabaseId   string `json:"database_id"`
	DatabaseName string `json:"database_name"`
	Region       string `json:"region"`
	Replicas     int    `json:"replicas"`
	Port         int    `json:"port"`
	CreationTime int64  `json:"creation_time"`
	Password     string `json:"password,omitempty"`
	User         string `json:"customer_id"`
	Endpoint     string `json:"endpoint"`
	Tls          bool   `json:"tls"`
	Consistent   bool   `json:"consistent"`
	MultiZone    bool   `json:"multi_zone"`
	RestToken    string `json:"rest_token,omitempty"`
}

type CreateDatabaseRequest struct {
	Region       string `json:"region"`
	DatabaseName string `json:"database_name"`
	Tls          bool   `json:"tls"`
	Consistent   bool   `json:"consistent"`
	MultiZone    bool   `json:"multi_zone"`
}

type Cluster struct {
	ClusterId   string `json:"cluster_id"`
	ClusterName string `json:"name"`
	Region      string `json:"region"`

	Type                 string `json:"type"`
	MultiZone            bool   `json:"multi_zone"`
	TcpEndpoint          string `json:"tcp_endpoint"`
	RestEndpoint         string `json:"rest_endpoint"`
	State                string `json:"state"`
	Username             string `json:"username"`
	EncodedUsername      string `json:"encoded_username"`
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
