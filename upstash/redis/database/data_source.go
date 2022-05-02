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
			},
			"multizone": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When enabled database is highly available and deployed multi-zone",
			},
			"tls": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When enabled data is encrypted in transit",
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
		},
	}
}
