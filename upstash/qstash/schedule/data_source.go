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
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"not_before": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"body": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
		},
	}
}