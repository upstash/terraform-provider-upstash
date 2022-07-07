package credential

type Credential struct {
	CredentialId   string `json:"credential_id"`
	CredentialName string `json:"credential_name"`
	Topic          string `json:"topic"`
	Permissions    string `json:"permissions"`
	ClusterId      string `json:"cluster_id"`
	Username       string `json:"username"`
	CreationTime   int64  `json:"creation_time"`
	State          string `json:"state"`
	Password       string `json:"password"`
}

type CreateCredentialRequest struct {
	CredentialName string `json:"credential_name"`
	ClusterId      string `json:"cluster_id"`
	Topic          string `json:"topic"`
	Permissions    string `json:"permissions"`
}

// credential_id	ID of the kafka credential
// credential_name	Name of the kafka credential
// topic	Name of the kafka topic
// permissions	Permission scope given to the kafka credential
// cluster_id	ID of the kafka cluster
// username	Username to be used for the kafka credential
// creation_time	Creation time of the credential
// state	State of the credential(active or deleted)
// password	Password to be used in authenticating to the cluste
