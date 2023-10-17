package schedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceQstashSchedule() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceScheduleRead,
		Schema: map[string]*schema.Schema{
			"schedule_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Qstash Schedule ID for requested schedule",
			},
			"cron": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cron string for Qstash Schedule",
			},
			"destination": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"retries": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Retries for Qstash Schedule. Either Topic ID or valid URL",
			},
			"not_before": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Start time for Qstash Schedule",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time for Qstash Schedule",
			},
			"body": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Encoded body for Qstash Schedule",
			},
			"forward_headers": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Forward headers to your API",
			},
		},
	}
}
