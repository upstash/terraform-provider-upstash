package connector

type Connector struct {
	ConnectorId         string                 `json:"connector_id"`
	Name                string                 `json:"name"`
	CustomerId          string                 `json:"customer_id"`
	ClusterId           string                 `json:"cluster_id"`
	CreationTime        int64                  `json:"creation_time"`
	DeletionTime        int64                  `json:"deletion_time"`
	State               string                 `json:"state"`
	StateErrorMessage   string                 `json:"state_error_message"`
	ConnectorState      string                 `json:"connector_state"`
	Tasks               []KafkaConnectorTask   `json:"tasks"`
	Topics              []string               `json:"topics"`
	ConnectorClass      string                 `json:"connector_class"`
	Properties          map[string]interface{} `json:"properties,omitempty"`
	PropertiesEncrypted string                 `json:"properties_encrypted,omitempty"`
	EncodedUsername     string                 `json:"encoded_username,omitempty"`
	UserPassword        string                 `json:"user_password,omitempty"`
	TTL                 int64                  `json:"TTL,omitempty"`
}

type KafkaConnectorTask struct {
	Id    int    `json:"id"`
	State string `json:"state"`
	Trace string `json:"trace"`
}

type CreateConnectorRequest struct {
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

type ReconfigureConnectorRequest struct {
	Properties map[string]interface{}
}
