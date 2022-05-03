package topic

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func ResourceTopic() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"topic_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Cluster ID for created topic",
			},
			"topic_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the topic",
			},
			"partitions": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The number of partitions the topic will have",
			},
			"retention_time": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Retention time of messages in the topic",
			},
			"retention_size": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Retention size of the messages in the topic",
			},
			"max_message_size": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Max message size in the topic",
			},
			"cleanup_policy": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cleanup policy will be used in the topic(compact or delete)",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the cluster the topic will be deployed in",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Region of the kafka topic",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the kafka topic (active or deleted)",
			},
			"multizone": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether multizone replication is enabled",
			},
			"tcp_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "TCP Endpoint of the kafka topic",
			},
			"rest_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "REST Endpoint of the kafka topic",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Base64 encoded username to be used in authenticating to the cluster",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password to be used in authenticating to the cluster",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time of the topic",
			},
		},
	}
}
