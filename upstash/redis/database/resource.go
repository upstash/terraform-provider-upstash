package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceDatabase() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDatabaseCreate,
		ReadContext:   resourceDatabaseRead,
		UpdateContext: resourceDatabaseUpdate,
		DeleteContext: resourceDatabaseDelete,
		Schema: map[string]*schema.Schema{
			"database_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Database ID for created database",
			},
			"database_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the database",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "region of the database. Possible values are: " +
					"\"global\", \"eu-west-1\", \"us-east-1\", \"us-west-1\", \"ap-northeast-1\" , \"eu-central1\"",
			},
			"endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Database URL for connection",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password of the database",
			},
			"consistent": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "When enabled, all writes are synchronously persisted to the disk.",
				Deprecated:  "Consistent option is deprecated.",
			},
			"multizone": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "When enabled, database becomes highly available and is deployed in multiple zones. (If changed to false from true, results in deletion and recreation of the resource)",
			},
			"tls": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "When enabled, data is encrypted in transit. (If changed to false from true, results in deletion and recreation of the resource)",
				Deprecated:  "TLS option is deprecated. TLS will always be enabled. If you have a DB without tls enabled, run the same configuration with tls=true to enable it.",
			},
			"port": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port of the endpoint",
			},
			"rest_token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Rest Token for the database.",
			},
			"read_only_rest_token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rest Token for the database.",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time of the database",
			},
			"database_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the database",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the database",
			},
			"user_email": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User email for the database",
			},
			"db_max_clients": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max clients for the database",
			},
			"db_max_request_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max request size for the database",
			},
			"db_disk_threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Disk threshold for the database",
			},
			"db_max_entry_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max entry size for the database",
			},
			"db_memory_threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Memory threshold for the database",
			},
			"db_daily_bandwidth_limit": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Daily bandwidth limit for the database",
			},
			"db_max_commands_per_second": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max commands per second for the database",
			},
		},

		CustomizeDiff: customdiff.All(
			customdiff.ForceNewIfChange("multizone", func(ctx context.Context, old, new, meta interface{}) bool {
				return old.(bool) && !new.(bool)
			}),
			customdiff.ForceNewIfChange("consistent", func(ctx context.Context, old, new, meta interface{}) bool {
				return old.(bool) && !new.(bool)
			}),
		),
	}
}
