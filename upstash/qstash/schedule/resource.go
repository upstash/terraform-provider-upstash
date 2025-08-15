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
			"schedule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Schedule ID for requested schedule",
			},
			"body": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Body to send for the POST request in string format. Needs escaping (\\) double quotes.",
			},
			"cron": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cron string for Qstash Schedule",
			},
			"destination": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Destination for Qstash Schedule. Either Topic ID or valid URL",
			},
			"content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Content type for Qstash Scheduling.",
			},
			"deduplication_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Deduplication ID for Qstash Scheduling.",
			},
			"content_based_deduplication": {
				Type:        schema.TypeBool,
				ForceNew:    true,
				Optional:    true,
				Description: "Content Based Deduplication (bool) for Qstash Scheduling.",
			},
			"not_before": {
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Description: "Start time for Qstash Scheduling.",
			},
			"delay": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Delay for Qstash Schedule.",
			},
			"callback": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Callback URL for Qstash Schedule.",
			},
			"retries": {
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Default:     3,
				Description: "Retries for Qstash Schedule requests.",
			},
			"created_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time for Qstash Schedule.",
			},
			"forward_headers": {
				Type:     schema.TypeMap,
				ForceNew: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Forward headers to your API",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
