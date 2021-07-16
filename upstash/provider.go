package upstash

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPSTASH_EMAIL", nil),
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("UPSTASH_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"upstash_database": resourceDatabase(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"upstash_database": dataSourceDatabase(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("email").(string)
	password := d.Get("api_key").(string)

	c := NewUpstashClient(username, password)
	return c, nil
}
