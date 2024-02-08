package database

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceDatabase() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceDatabaseRead,
		Schema: map[string]*schema.Schema{
			"database_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Database ID for requested database",
			},
			"database_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the database",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
				Computed:    true,
				Description: "When enabled database runs in Consistency Mode",
				Deprecated:  "Consistent option is deprecated.",
			},
			"multizone": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When enabled database is highly available and deployed multi-zone",
				Deprecated:  "Multizone option is deprecated. It is enabled by default for paid databases.",
			},
			"tls": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When enabled data is encrypted in transit",
			},
			"eviction": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable eviction, to evict keys when your database reaches the max size",
			},
			"auto_scale": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Upgrade to higher plans automatically when it hits quotas",
			},
			"primary_region": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Primary region for the database",
			},
			"read_regions": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Read regions for the database",
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
				Sensitive:   true,
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
	}
}
