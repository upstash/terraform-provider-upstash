package topic

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceTopic() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceRead,
		Schema: map[string]*schema.Schema{
			"topic_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Topic ID for requested kafka topic",
			},
			"topic_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the kafka topic",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the cluster this topic belongs to",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Description: "Region of the kafka topic. Possible values (may change) are: " +
					"\"eu-west-1\", \"us-east-1\"",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time of the topic",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the topic (active or deleted)",
			},
			"partitions": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of partitions the topic has",
			},
			"multizone": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether multizone replication is enabled",
			},
			"tcp_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "TCP Endpoint of the topic",
			},
			"rest_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "REST Endpoint of the topic",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username to be used in authenticating to the cluster",
			},
			"encoded_username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Base64 encoded username to be used in rest communication",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password to be used in authenticating to the cluster",
			},
			"cleanup_policy": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cleanup policy will be used in the topic(compact or delete)",
			},
			"retention_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Retention Size of the topic",
			},
			"retention_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Retention Time of the topic",
			},
			"max_message_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Max Message Size for the topic",
			},
		},
	}
}
