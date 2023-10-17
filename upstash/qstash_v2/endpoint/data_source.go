package endpoint

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceQstashEndpoint() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceEndpointRead,
		Schema: map[string]*schema.Schema{
			"topic_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Topic Id that the endpoint is added to",
			},
			"endpoint_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Endpoint ID",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the endpoint",
			},
		},
	}
}
