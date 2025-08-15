package database

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceDatabase() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceDatabaseRead,
		Schema: map[string]*schema.Schema{
			"database_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Database ID for requested database",
			},
			"database_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the database",
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "region of the database. Possible values are: " +
					"\"global\", \"us-central1\"",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Database URL for connection",
			},
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password of the database",
			},
			"consistent": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When enabled database runs in Consistency Mode",
				Deprecated:  "Consistent option is deprecated.",
			},
			"multizone": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When enabled database is highly available and deployed multi-zone",
				Deprecated:  "Multizone option is deprecated. It is enabled by default for paid databases.",
			},
			"tls": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When enabled, data is encrypted in transit.",
				Deprecated:  "TLS option is deprecated. It's enabled by default for all databases.",
			},
			"eviction": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable eviction, to evict keys when your database reaches the max size",
			},
			"auto_upgrade": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Upgrade to higher plans automatically when it hits quotas",
			},
			"budget": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Budget for the database. It is used to limit the cost of the database. If the budget is reached, the database will be throttled until the next month.",
			},
			"primary_region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Primary region for the database",
			},
			"read_regions": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Read regions for the database",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port of the endpoint",
			},
			"rest_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Rest Token for the database.",
			},
			"read_only_rest_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Rest Token for the database.",
			},
			"creation_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time of the database",
			},
			"database_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the database",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the database",
			},
			"user_email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User email for the database",
			},
			"db_max_clients": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max clients for the database",
			},
			"db_max_request_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max request size for the database",
			},
			"db_disk_threshold": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Disk threshold for the database",
			},
			"db_max_entry_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max entry size for the database",
			},
			"db_memory_threshold": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Memory threshold for the database",
			},
			"db_daily_bandwidth_limit": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Daily bandwidth limit for the database",
			},
			"db_max_commands_per_second": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max commands per second for the database",
			},
		},
	}
}
