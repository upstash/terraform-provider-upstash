package topic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceQstashTopic() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceTopicRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Qstash Topic",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time for Qstash Topic.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last Update time for Qstash Topic.",
			},
			"endpoints": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Endpoints for the Qstash Topic",
			},
		},
	}
}
