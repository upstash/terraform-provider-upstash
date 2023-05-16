package connector

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnector() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"connector_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Connector ID for created connector",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the connector",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the connector",
			},
			"properties": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Properties that the connector will have",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation of the connector",
			},
			"running_state": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Running state of the connector. Can be either 'paused', 'running' or 'restart'",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != "paused" && val != "running" && val != "restart" {
						errs = append(errs, fmt.Errorf("running_state field can only take the values: [`paused`, `running`, `restart`]"))
					}
					return
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
