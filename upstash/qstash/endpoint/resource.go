package endpoint

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func ResourceQstashEndpoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEndpointCreate,
		ReadContext:   resourceEndpointRead,
		DeleteContext: resourceEndpointDelete,
		Schema: map[string]*schema.Schema{
			"topic_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Topic Id that the endpoint is added to",
			},
			"endpoint_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Endpoint ID",
			},
			"topic_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Topic Name for Endpoint",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "URL of the endpoint",
			},
		},
	}
}
