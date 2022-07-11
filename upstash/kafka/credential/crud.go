package credential

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceCredentialCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	credential, err := createCredential(c, CreateCredentialRequest{
		CredentialName: data.Get("credential_name").(string),
		ClusterId:      data.Get("cluster_id").(string),
		Topic:          data.Get("topic").(string),
		Permissions:    data.Get("permissions").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-kafka-credential-" + credential.CredentialId)
	data.Set("credential_id", credential.CredentialId)
	return resourceCredentialRead(ctx, data, m)
}

func resourceCredentialDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	credentialId := data.Get("credential_id").(string)
	err := deleteCredential(c, credentialId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceCredentialRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	credentialId := data.Get("credential_id").(string)
	credential, err := getCredential(c, credentialId)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-kafka-credential-" + credential.CredentialId)

	mapping := map[string]interface{}{
		"credential_id":   credential.CredentialId,
		"credential_name": credential.CredentialName,
		"topic":           credential.Topic,
		"permissions":     credential.Permissions,
		"cluster_id":      credential.ClusterId,
		"username":        credential.Username,
		"creation_time":   credential.CreationTime,
		"state":           credential.State,
		"password":        credential.Password,
	}
	return utils.SetAndCheckErrors(data, mapping)
}
