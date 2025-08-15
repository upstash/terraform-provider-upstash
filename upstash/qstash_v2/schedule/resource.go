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
			"created_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time for Qstash Schedule.",
			},
			"schedule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Qstash Schedule ID for requested schedule",
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
				Description: "Destination for Qstash Schedule. Either Topic Name or valid URL",
			},
			"method": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Method of Http Request on QStash",
			},
			"header": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Headers for the QStash schedule",
			},
			"body": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Body to send for the POST request in string format. Needs escaping (\\) double quotes.",
			},
			"retries": {
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Default:     3,
				Description: "Retries for Qstash Schedule requests.",
			},
			"delay": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Default:     "0",
				Description: "Delay for Qstash Schedule. Delay should be given in seconds",
			},
			"callback": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Callback URL for Qstash Schedule.",
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
			"content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Content type for Qstash Scheduling.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
