package topic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/qstash/endpoint"
)

func ResourceQstashTopic() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTopicCreate,
		ReadContext:   resourceTopicRead,
		DeleteContext: resourceTopicDelete,
		Schema: map[string]*schema.Schema{
			"topic_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Topic ID for requested topic",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the Qstash Topic",
			},
			// TODO: Use resource after generating endpoints resource
			"endpoints": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        endpoint.ResourceQstashEndpoint(),
				Description: "Members of the team. Email addresses are given as the keys with their roles as the values.",
			},
		},
	}
}
