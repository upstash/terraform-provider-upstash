package cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceClusterCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	cluster, err := createCluster(c, CreateClusterRequest{
		ClusterName: data.Get("cluster_name").(string),
		Region:      data.Get("region").(string),
		MultiZone:   data.Get("multizone").(bool),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-cluster-" + cluster.ClusterId)
	data.Set("cluster_id", cluster.ClusterId)
	return resourceClusterRead(ctx, data, m)

}

// Only name can change in v2 api.
func resourceClusterUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	clusterId := data.Get("cluster_id").(string)
	if data.HasChange("cluster_name") {
		if err := renameCluster(c, clusterId, data.Get("cluster_name").(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceClusterRead(ctx, data, m)
}

func resourceClusterDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	clusterId := data.Get("cluster_id").(string)
	err := deleteCluster(c, clusterId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceClusterRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	clusterId := data.Get("cluster_id").(string)
	cluster, err := getCluster(c, clusterId)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-cluster-" + cluster.ClusterId)

	mapping := map[string]interface{}{
		"cluster_id":              cluster.ClusterId,
		"cluster_name":            cluster.ClusterName,
		"region":                  cluster.Region,
		"type":                    cluster.Type,
		"multizone":               cluster.MultiZone,
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

	return utils.SetAndCheckErrors(data, mapping)
}
