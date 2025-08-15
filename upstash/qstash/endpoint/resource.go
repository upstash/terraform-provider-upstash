package endpoint

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func ResourceQstashEndpoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEndpointCreate,
		ReadContext:   resourceEndpointRead,
		UpdateContext: resourceEndpointUpdate,
		DeleteContext: resourceEndpointDelete,
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Topic Id that the endpoint is added to",
			},
			"endpoint_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Endpoint ID",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Topic Name for Endpoint",
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL of the endpoint",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
