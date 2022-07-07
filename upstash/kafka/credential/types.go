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
