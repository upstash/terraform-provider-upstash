package topic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceQstashTopic() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceTopicRead,
		Schema: map[string]*schema.Schema{
			"topic_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Qstash Topic ID for requested topic",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Qstash Topic",
			},
			"endpoints": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				Description: "Endpoints for the Qstash Topic",
			},
		},
	}
}
