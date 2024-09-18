package cluster

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceClusterRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Cluster ID for requested cluster",
			},
			"cluster_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the cluster",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Description: "Region of the cluster. Possible values (may change) are: " +
					"\"eu-west-1\", \"us-east-1\"",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the cluster",
			},
			"multizone": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether multizone replication is enabled",
			},
			"tcp_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "TCP Endpoint of the cluster",
			},
			"rest_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "REST Endpoint of the cluster",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current state of the cluster (active or deleted)",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Base64 encoded username for the cluster",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password for the cluster",
			},
			"max_retention_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Retention Size of the cluster",
			},
			"max_retention_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Retention Time of the cluster",
			},
			"max_messages_per_second": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Messages Per Second for the cluster",
			},
			"max_message_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Message Size for the cluster",
			},
			"max_partitions": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Partitions for the cluster",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time of the cluster",
			},
		},
		DeprecationMessage: "Upstash Kafka service is deprecated. It's no longer possible to create new clusters, but existing clusters will be continued until March 2025. For further information regarding this issue, see https://upstash.com/blog/workflow-kafka",
	}
}
