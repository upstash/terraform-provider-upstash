package connector

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	clusterId := data.Get("cluster_id").(string)
	connector, err := createConnector(c, clusterId, CreateConnectorRequest{
		Name:       data.Get("name").(string),
		Properties: data.Get("properties").(map[string]interface{}),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-kafka-connector-" + connector.ConnectorId)
	data.Set("connector_id", connector.ConnectorId)
	return resourceRead(ctx, data, m)
}

func resourceRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	connectorId := data.Get("connector_id").(string)
	if connectorId == "" {
		connectorId = data.Id()
	}
	connector, err := getConnector(c, connectorId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-kafka-connector-" + connector.ConnectorId)

	mapping := map[string]interface{}{
		"connector_id":  connector.ConnectorId,
		"name":          connector.Name,
		"cluster_id":    connector.ClusterId,
		"creation_time": connector.CreationTime,
		"properties":    connector.Properties,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	connectorId := data.Get("connector_id").(string)
	clusterId := data.Get("cluster_id").(string)
	if data.HasChange("properties") {
		err := reconfigureConnector(c, clusterId, connectorId, data.Get("properties").(map[string]interface{}))

		if err != nil {
			return diag.FromErr(err)
		}
	}
	if data.HasChange("running_state") {
		connectorState := data.Get("running_state").(string)
		var err error

		if connectorState == "paused" {
			err = pauseConnector(c, clusterId, connectorId)
		} else if connectorState == "running" {
			err = startConnector(c, clusterId, connectorId)
		} else if connectorState == "restart" {
			err = restartConnector(c, clusterId, connectorId)
		}
		if err != nil {
			data.Set("running_state", "")
			return diag.FromErr(err)
		}
	}
	return resourceRead(ctx, data, m)
}

func resourceDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	connectorId := data.Get("connector_id").(string)
	clusterId := data.Get("cluster_id").(string)
	err := deleteConnector(c, clusterId, connectorId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceReadForDataSource(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	connectorId := data.Get("connector_id").(string)
	connector, err := getConnector(c, connectorId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-kafka-connector-" + connector.ConnectorId)

	var tasks []map[string]interface{}
	for _, task := range connector.Tasks {
		tasks = append(tasks, map[string]interface{}{
			"id":    strconv.Itoa(task.Id),
			"state": task.State,
			"trace": task.Trace,
		})
	}

	mapping := map[string]interface{}{
		"connector_id":         connector.ConnectorId,
		"name":                 connector.Name,
		"cluster_id":           connector.ClusterId,
		"creation_time":        connector.CreationTime,
		"state":                connector.State,
		"state_error_message":  connector.StateErrorMessage,
		"connector_state":      connector.ConnectorState,
		"tasks":                tasks,
		"topics":               connector.Topics,
		"connector_class":      connector.ConnectorClass,
		"properties":           connector.Properties,
		"properties_encrypted": connector.PropertiesEncrypted,
		"encoded_username":     connector.EncodedUsername,
		"user_password":        connector.UserPassword,
		"ttl":                  connector.TTL,
	}

	return utils.SetAndCheckErrors(data, mapping)
}
