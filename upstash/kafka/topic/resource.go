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
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The number of partitions the topic will have",
			},
			"retention_time": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Retention time of messsages in the topic",
			},
			"retention_size": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Retention size of the messages in the topic",
			},
			"max_message_size": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Max message size in the topic",
			},
			"cleanup_policy": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cleanup policy will be used in the topic(compact or delete)",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the cluster the topic will be deployed in",
			},
		},
	}
}
