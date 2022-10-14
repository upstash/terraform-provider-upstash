package schedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceQstashSchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceScheduleCreate,
		ReadContext:   resourceScheduleRead,
		DeleteContext: resourceScheduleDelete,
		Schema: map[string]*schema.Schema{
			"schedule_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Schedule ID for requested schedule",
			},
			"body": &schema.Schema{
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Body to send for the POST request in string format. Needs escaping (\\) double quotes.",
			},
			"cron": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cron string for Qstash Schedule",
			},
			"destination": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"content_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"deduplication_id": &schema.Schema{
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"content_based_deduplication": &schema.Schema{
				Type:        schema.TypeBool,
				ForceNew:    true,
				Optional:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"not_before": &schema.Schema{
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"delay": &schema.Schema{
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"retries": &schema.Schema{
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
		},
	}
}
