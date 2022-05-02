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
				ForceNew:    true,
				Description: "When enabled, all writes are synchronously persisted to the disk.",
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
				Default:     false,
				Description: "When enabled, data is encrypted in transit. (If changed to false from true, results in deletion and recreation of the resource)",
			},
			"port": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port of the endpoint",
			},
		},

		CustomizeDiff: customdiff.All(
			customdiff.ForceNewIfChange("multizone", func(ctx context.Context, old, new, meta interface{}) bool {
				return old.(bool) && !new.(bool)
			}),
			customdiff.ForceNewIfChange("tls", func(ctx context.Context, old, new, meta interface{}) bool {
				return old.(bool) && !new.(bool)
			}),
		),
	}
}
