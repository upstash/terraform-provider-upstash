package upstash

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDatabase() *schema.Resource {
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
					"\"eu-west-1\", \"us-east-1\", \"us-west-1\", \"ap-northeast-1\" , \"eu-central1\"",
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
				Description: "When enabled database runs in Consistency Mode",
			},
			"multi_zone": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    false,
				Description: "When enabled database is highly available and deployed multi-zone",
			},
			"tls": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    false,
				Description: "When enabled data is encrypted in transit",
			},
		},
	}
}

func resourceDatabaseUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	databaseId := data.Get("database_id").(string)
	if data.HasChange("multi_zone") {
		if err := c.EnableMultiZone(databaseId, data.Get("multi_zone").(bool)); err != nil {
			return diag.FromErr(err)
		}
	}
	// I think this creates a problem. Doesnt give tls as a parameter for the client to use.
	if data.HasChange("tls") {
		if err := c.EnableTLS(databaseId); err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceDatabaseRead(ctx, data, m)
}

func resourceDatabaseDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	databaseId := data.Get("database_id").(string)
	err := c.DeleteDatabase(databaseId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceDatabaseRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	databaseId := data.Get("database_id").(string)
	database, err := c.GetDatabase(databaseId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = data.Set("database_id", database.DatabaseId); err != nil {
		return diag.FromErr(err)
	}

	if err = data.Set("database_name", database.DatabaseName); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("region", database.Region); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("endpoint", database.Endpoint); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("password", database.Password); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("consistent", database.Consistent); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("multi_zone", database.MultiZone); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("tls", database.Tls); err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-database-" + database.DatabaseId)
	return nil
}

func resourceDatabaseCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	database, err := c.CreateDatabase(CreateDatabaseRequest{
		Region:       data.Get("region").(string),
		DatabaseName: data.Get("database_name").(string),
		Tls:          data.Get("tls").(bool),
		Consistent:   data.Get("consistent").(bool),
		MultiZone:    data.Get("multi_zone").(bool),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-database-" + database.DatabaseId)
	data.Set("database_id", database.DatabaseId)
	return resourceDatabaseRead(ctx, data, m)
}
