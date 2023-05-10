package credential

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCredential() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCredentialCreate,
		ReadContext:   resourceCredentialRead,
		DeleteContext: resourceCredentialDelete,
		Schema: map[string]*schema.Schema{
			"credential_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique ID of the kafka credential",
			},
			"credential_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the kafka credential",
			},
			"topic": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the kafka topic",
			},
			"permissions": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Permission scope given to the kafka credential",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != "ALL" && val != "CONSUME" && val != "PRODUCE" {
						errs = append(errs, fmt.Errorf("permissions field can only take the values: [`ALL`, `PRODUCE`, `CONSUME`]"))
						// errs = append(errs, fmt.Errorf("Owner of the api key should be given the role of owner"))
					}
					return
				},
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the kafka cluster",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username to be used for the kafka credential",
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
				Sensitive:   true,
				Description: "Password to be used in authenticating to the cluster",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
