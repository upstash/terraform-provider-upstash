package topic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceTopicRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicId := data.Get("topic_id").(string)
	topic, err := getTopic(c, topicId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-topic-" + topic.TopicId)

	mapping := map[string]interface{}{
		"name":      topic.Name,
		"topic_id":  topic.TopicId,
		"endpoints": topic.Endpoints,
	}

	return utils.SetAndCheckErrors(data, mapping)

}

func resourceTopicCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topic, err := createTopic(c, createQstashTopicRequest{
		Name: data.Get("name").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-qstash-topic-" + topic.TopicId)
	data.Set("topic_id", topic.TopicId)
	return resourceTopicRead(ctx, data, m)
}

func resourceTopicDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicId := data.Get("topic_id").(string)
	err := deleteTopic(c, topicId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
