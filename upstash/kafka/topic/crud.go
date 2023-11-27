package topic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topic, err := createTopic(c, CreateTopicRequest{
		TopicName:      data.Get("topic_name").(string),
		Partitions:     data.Get("partitions").(int),
		RetentionTime:  data.Get("retention_time").(int),
		RetentionSize:  data.Get("retention_size").(int),
		MaxMessageSize: data.Get("max_message_size").(int),
		CleanupPolicy:  data.Get("cleanup_policy").(string),
		ClusterId:      data.Get("cluster_id").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-kafka-topic-" + topic.TopicId)
	data.Set("topic_id", topic.TopicId)
	return resourceRead(ctx, data, m)
}

func resourceRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicId := data.Get("topic_id").(string)
	if topicId == "" {
		topicId = data.Id()
	}

	topic, err := getTopic(c, topicId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-kafka-topic-" + topic.TopicId)

	mapping := map[string]interface{}{
		"topic_id":         topic.TopicId,
		"topic_name":       topic.TopicName,
		"partitions":       topic.Partitions,
		"retention_time":   topic.MaxRetentionTime,
		"retention_size":   topic.MaxRetentionSize,
		"max_message_size": topic.MaxMessageSize,
		"cleanup_policy":   topic.CleanupPolicy,
		"cluster_id":       topic.ClusterId,
		"region":           topic.Region,
		"state":            topic.State,
		"multizone":        topic.MultiZone,
		"tcp_endpoint":     topic.TcpEndpoint,
		"rest_endpoint":    topic.RestEndpoint,
		"username":         topic.Username,
		"password":         topic.Password,
		"creation_time":    topic.CreationTime,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicId := data.Get("topic_id").(string)
	if data.HasChange("retention_time") || data.HasChange("retention_size") || data.HasChange("max_message_size") || data.HasChange("partitions") {
		err := reconfigureKafkaTopic(c, topicId, ReconfigureTopic{
			RetentionTime:  data.Get("retention_time").(int),
			RetentionSize:  data.Get("retention_size").(int),
			MaxMessageSize: data.Get("max_message_size").(int),
			Partitions:     data.Get("partitions").(int),
		})

		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceRead(ctx, data, m)
}

func resourceDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicId := data.Get("topic_id").(string)
	err := deleteTopic(c, topicId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
