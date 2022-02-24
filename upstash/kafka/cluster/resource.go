package cluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		// rename these, like in team, also redis/db ones
		CreateContext: resourceClusterCreate,
		ReadContext:   resourceClusterRead,
		UpdateContext: resourceClusterUpdate,
		DeleteContext: resourceClusterDelete,
		Schema: map[string]*schema.Schema{

			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Cluster ID for created cluster",
			},
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// Not sure about this
				// ForceNew:    true,
				Description: "Name of the cluster",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "region of the cluster. Possible values (may change) are: " +
					"\"eu-west-1\", \"us-east-1\"",
			},

			// not implemented yet
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the cluster",
			},
			"multizone": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Whether cluster has multizone attribute",
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
				Description: "State, where the cluster is originated",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username for the cluster",
			},
			"encoded_username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Encoded Username for the cluster",
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
		},
	}

}
