package upstash

import (
	"context"

	"github.com/upstash/terraform-provider-upstash/upstash/vector/index"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	qstashEndpoint "github.com/upstash/terraform-provider-upstash/upstash/qstash/endpoint"
	qstashSchedule "github.com/upstash/terraform-provider-upstash/upstash/qstash/schedule"
	qstashTopic "github.com/upstash/terraform-provider-upstash/upstash/qstash/topic"
	qstashScheduleV2 "github.com/upstash/terraform-provider-upstash/upstash/qstash_v2/schedule"
	qstashTopicV2 "github.com/upstash/terraform-provider-upstash/upstash/qstash_v2/topic"

	"github.com/upstash/terraform-provider-upstash/upstash/redis/database"
	"github.com/upstash/terraform-provider-upstash/upstash/team"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPSTASH_EMAIL", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("UPSTASH_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"upstash_redis_database":     database.ResourceDatabase(),
			"upstash_vector_index":       index.ResourceIndex(),
			"upstash_team":               team.ResourceTeam(),
			"upstash_qstash_topic":       qstashTopic.ResourceQstashTopic(),
			"upstash_qstash_endpoint":    qstashEndpoint.ResourceQstashEndpoint(),
			"upstash_qstash_schedule":    qstashSchedule.ResourceQstashSchedule(),
			"upstash_qstash_topic_v2":    qstashTopicV2.ResourceQstashTopic(),
			"upstash_qstash_schedule_v2": qstashScheduleV2.ResourceQstashSchedule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"upstash_redis_database_data":     database.DataSourceDatabase(),
			"upstash_vector_index_data":       index.DataResourceIndex(),
			"upstash_team_data":               team.DataSourceTeam(),
			"upstash_qstash_topic_data":       qstashTopic.DataSourceQstashTopic(),
			"upstash_qstash_endpoint_data":    qstashEndpoint.DataSourceQstashEndpoint(),
			"upstash_qstash_schedule_data":    qstashSchedule.DataSourceQstashSchedule(),
			"upstash_qstash_topic_v2_data":    qstashTopicV2.DataSourceQstashTopic(),
			"upstash_qstash_schedule_v2_data": qstashScheduleV2.DataSourceQstashSchedule(),
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
