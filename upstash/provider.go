package upstash

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/kafka/cluster"
	"github.com/upstash/terraform-provider-upstash/upstash/kafka/credential"
	"github.com/upstash/terraform-provider-upstash/upstash/kafka/topic"
	qstashEndpoint "github.com/upstash/terraform-provider-upstash/upstash/qstash/endpoint"
	qstashTopic "github.com/upstash/terraform-provider-upstash/upstash/qstash/topic"

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
			"upstash_redis_database":   database.ResourceDatabase(),
			"upstash_kafka_cluster":    cluster.ResourceCluster(),
			"upstash_kafka_topic":      topic.ResourceTopic(),
			"upstash_kafka_credential": credential.ResourceCredential(),
			"upstash_team":             team.ResourceTeam(),
			"upstash_qstash_topic":     qstashTopic.ResourceQstashTopic(),
			"upstash_qstash_endpoint":  qstashEndpoint.ResourceQstashEndpoint(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"upstash_redis_database_data":   database.DataSourceDatabase(),
			"upstash_kafka_cluster_data":    cluster.DataSourceCluster(),
			"upstash_kafka_topic_data":      topic.DataSourceTopic(),
			"upstash_team_data":             team.DataSourceTeam(),
			"upstash_kafka_credential_data": credential.DataSourceCredential(),
			"upstash_qstash_topic_data":     qstashTopic.DataSourceQstashTopic(),
			"upstash_qstash_endpoint_data":  qstashEndpoint.DataSourceQstashEndpoint(),
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
