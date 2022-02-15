package upstash

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/kafka/cluster"
	"github.com/upstash/terraform-provider-upstash/upstash/kafka/topic"
	"github.com/upstash/terraform-provider-upstash/upstash/redis/database"
	"github.com/upstash/terraform-provider-upstash/upstash/team"

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
			"upstash_redis_database": database.ResourceDatabase(),
			"upstash_kafka_cluster":  cluster.ResourceCluster(),
			"upstash_kafka_topic":    topic.ResourceTopic(),
			"upstash_team":           team.ResourceTeam(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"upstash_redis_database_data": database.DataSourceDatabase(),
			"upstash_kafka_cluster_data":  cluster.DataSourceCluster(),
			"upstash_kafka_topic_data":    topic.DataSourceTopic(),
			"upstash_team_data":           team.DataSourceTeam(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("email").(string)
	password := d.Get("api_key").(string)

	c := client.NewUpstashClient(username, password)
	return c, nil
}
