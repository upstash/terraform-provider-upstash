package schedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceQstashSchedule() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceScheduleRead,
		Schema: map[string]*schema.Schema{
			"created_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time for Qstash Schedule.",
			},
			"schedule_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Qstash Schedule ID for requested schedule",
			},
			"cron": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cron string for Qstash Schedule",
			},
			"destination": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination for Qstash Schedule. Either Topic Name or valid URL",
			},
			"method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Method of Http Request on QStash",
			},
			"header": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Headers for the QStash schedule",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Body to send for the POST request in string format. Needs escaping (\\) double quotes.",
			},
			"retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3,
				Description: "Retries for Qstash Schedule requests.",
			},
			"delay": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Delay for Qstash Schedule.",
			},
			"callback": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Callback URL for Qstash Schedule.",
			},
		},
	}
}
