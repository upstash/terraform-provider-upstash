package credential

import (
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func createCredential(c *client.UpstashClient, body CreateCredentialRequest) (credential Credential, err error) {
	resp, err := c.SendPostRequest("/v2/kafka/credential", body, "Create Kafka Credential", false)

	if err != nil {
		return credential, err
	}

	err = resp.ToJSON(&credential)
	return credential, err
}

func deleteCredential(c *client.UpstashClient, credentialId string) (err error) {
	return c.SendDeleteRequest("/v2/kafka/credential/"+credentialId, nil, "Delete Kafka Credential", false)
}

func getCredential(c *client.UpstashClient, credentialId string) (credential Credential, err error) {
	resp, err := c.SendGetRequest("/v2/kafka/credentials", "Get Kafka Cluster", false)
	var credentials []Credential
	if err != nil {
		return credential, err
	}
	err = resp.ToJSON(&credentials)
	if err != nil {
		return credential, err
	}

	for _, cd := range credentials {
		if cd.CredentialId == credentialId {
			credential = cd
		}
	}

	return credential, err
}
