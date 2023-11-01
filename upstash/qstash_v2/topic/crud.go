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
	topicName := data.Get("name").(string)

	if topicName == "" {
		topicName = data.Id()
	}

	topic, err := getTopic(c, topicName)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-topic-" + topic.Name)

	var endpointMap []string
	for _, val := range topic.Endpoints {
		endpointMap = append(endpointMap, val.Url)
	}

	mapping := map[string]interface{}{
		"name":       topic.Name,
		"created_at": topic.UpdatedAt,
		"updated_at": topic.UpdatedAt,
		"endpoints":  endpointMap,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceTopicCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicName := data.Get("name").(string)

	var endpoints []QStashEndpoint
	for _, v := range (data.Get("endpoints").(*schema.Set)).List() {
		if v != nil {
			endpoints = append(endpoints, QStashEndpoint{Url: v.(string)})
		}
	}

	err := createTopic(c, topicName, UpdateQStashTopicEndpoints{Endpoints: endpoints})
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-topic-" + topicName)
	data.Set("name", topicName)
	return resourceTopicRead(ctx, data, m)
}

func resourceTopicDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicName := data.Get("name").(string)
	err := deleteTopic(c, topicName)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceTopicUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	topicName := data.Get("name").(string)

	if data.HasChange("endpoints") {
		a, b := data.GetChange("endpoints")
		old := a.(*schema.Set).List()
		new := b.(*schema.Set).List()

		var endpointsToRemove []QStashEndpoint
		var endpointsToAdd []QStashEndpoint

		if len(new) == 0 {
			return diag.Errorf("At least 1 Url is required for a topic.")
		}

		for _, oldEndpoint := range old {
			needsRemoval := true
			for _, newEndpoint := range new {
				if newEndpoint.(string) == oldEndpoint.(string) {
					needsRemoval = false
					break
				}
			}
			if needsRemoval {
				endpointsToRemove = append(endpointsToRemove, QStashEndpoint{Url: oldEndpoint.(string)})
			}
		}

		for _, newEndpoint := range new {
			needsAdding := true
			for _, oldEndpoint := range old {
				if newEndpoint.(string) == oldEndpoint.(string) {
					needsAdding = false
					break
				}
			}
			if needsAdding {
				endpointsToAdd = append(endpointsToAdd, QStashEndpoint{Url: newEndpoint.(string)})
			}
		}

		if len(endpointsToAdd) > 0 {
			if err := addEndpointsToTopic(c, topicName, UpdateQStashTopicEndpoints{Endpoints: endpointsToAdd}); err != nil {
				return diag.FromErr((err))
			}
		}

		if len(endpointsToRemove) > 0 {
			if err := deleteEndpointsFromTopic(c, topicName, UpdateQStashTopicEndpoints{Endpoints: endpointsToRemove}); err != nil {
				return diag.FromErr((err))
			}
		}

	}

	return resourceTopicRead(ctx, data, m)
}
