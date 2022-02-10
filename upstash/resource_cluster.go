package upstash

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
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
					"\"eu-west-1\", \"us-east-1\", \"us-west-1\", \"ap-northeast-1\" , \"eu-central1\"",
			},

			// not implemented yet
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the cluster",
			},
			"multi_zone": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Whether cluster has multi_zone attribute",
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

// Only name can change in v2 api.
func resourceClusterUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	clusterId := data.Get("cluster_id").(string)
	if data.HasChange("cluster_name") {
		if err := c.RenameCluster(clusterId, data.Get("cluster_name").(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceClusterRead(ctx, data, m)
}

func resourceClusterDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	clusterId := data.Get("cluster_id").(string)
	err := c.DeleteCluster(clusterId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceClusterRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	clusterId := data.Get("cluster_id").(string)
	cluster, err := c.GetCluster(clusterId)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-cluster-" + cluster.ClusterId)

	mapping := map[string]interface{}{
		"cluster_id":              cluster.ClusterId,
		"cluster_name":            cluster.ClusterName,
		"region":                  cluster.Region,
		"type":                    cluster.Type,
		"multi_zone":              cluster.MultiZone,
		"tcp_endpoint":            cluster.TcpEndpoint,
		"rest_endpoint":           cluster.RestEndpoint,
		"state":                   cluster.State,
		"username":                cluster.Username,
		"encoded_username":        cluster.EncodedUsername,
		"password":                cluster.Password,
		"max_retention_size":      cluster.MaxRetentionSize,
		"max_retention_time":      cluster.MaxRetentionTime,
		"max_messages_per_second": cluster.MaxMessagesPerSecond,
		"max_message_size":        cluster.MaxMessageSize,
		"max_partitions":          cluster.MaxPartitions,
	}

	return setAndCheckErrors(data, mapping)

	// if err = data.Set("cluster_id", cluster.ClusterId); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("cluster_name", cluster.ClusterName); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("region", cluster.Region); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("type", cluster.Type); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("multi_zone", cluster.MultiZone); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("tcp_endpoint", cluster.TcpEndpoint); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("rest_endpoint", cluster.RestEndpoint); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("state", cluster.State); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("username", cluster.Username); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("encoded_username", cluster.EncodedUsername); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("password", cluster.Password); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("max_retention_size", cluster.MaxRetentionSize); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("max_retention_time", cluster.MaxRetentionTime); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("max_messages_per_second", cluster.MaxMessagesPerSecond); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("max_message_size", cluster.MaxMessageSize); err != nil {
	// 	return diag.FromErr(err)
	// }
	// if err = data.Set("max_retention_size", cluster.MaxRetentionSize); err != nil {
	// 	return diag.FromErr(err)
	// }

}

func resourceClusterCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*UpstashClient)
	cluster, err := c.CreateCluster(CreateClusterRequest{
		ClusterName: data.Get("cluster_name").(string),
		Region:      data.Get("region").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-cluster-" + cluster.ClusterId)
	data.Set("cluster_id", cluster.ClusterId)
	return resourceClusterRead(ctx, data, m)

}

func setAndCheckErrors(data *schema.ResourceData, mapping map[string]interface{}) diag.Diagnostics {
	for str, value := range mapping {
		if err := data.Set(str, value); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}
