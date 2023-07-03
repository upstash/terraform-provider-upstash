package endpoint

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceEndpointRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	endpointId := data.Get("endpoint_id").(string)
	if endpointId == "" {
		endpointId = data.Id()
	}

	endpoint, err := getEndpoint(c, endpointId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-endpoint-" + endpoint.EndpointId)

	mapping := map[string]interface{}{
		"endpoint_id": endpoint.EndpointId,
		"topic_id":    endpoint.TopicId,
		"url":         endpoint.Url,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceEndpointCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	endpoint, err := createEndpoint(c, createQstashEndpointRequest{
		TopicName: data.Get("topic_name").(string),
		TopicId:   data.Get("topic_id").(string),
		Url:       data.Get("url").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-qstash-endpoint-" + endpoint.EndpointId)
	data.Set("endpoint_id", endpoint.EndpointId)
	return resourceEndpointRead(ctx, data, m)
}

func resourceEndpointDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	endpointId := data.Get("endpoint_id").(string)
	err := deleteEndpoint(c, endpointId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceEndpointUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	endpointId := data.Get("endpoint_id").(string)
	if data.HasChange("url") {
		err := updateEndpoint(c, endpointId, UpdateQstashEndpoint{
			Url: data.Get("url").(string),
		})

		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceEndpointRead(ctx, data, m)
}
