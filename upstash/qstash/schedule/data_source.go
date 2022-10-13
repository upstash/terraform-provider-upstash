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
		},
	}
}
