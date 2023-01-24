package credential

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceCredential() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceCredentialRead,
		Schema: map[string]*schema.Schema{
			"credential_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique ID of the kafka credential",
			},
			"credential_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the kafka credential",
			},
			"topic": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the kafka topic",
			},
			"permissions": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Permission scope given to the kafka credential",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the kafka cluster",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username to be used for the kafka credential",
			},
			"decoded_username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Decoded Username to be used for the kafka credential",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time of the credential",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the credential(active or deleted)",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password to be used in authenticating to the cluster",
			},
		},
	}
}
