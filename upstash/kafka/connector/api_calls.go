package connector

import (
	"fmt"

	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func createConnector(c *client.UpstashClient, body CreateConnectorRequest) (connector Connector, err error) {

	resp, err := c.SendPostRequest("/v2/kafka/connector", body, "Create Kafka Connector", false)

	if err != nil {
		return connector, err
	}

	err = resp.ToJSON(&connector)
	return connector, err

}

func reconfigureConnector(c *client.UpstashClient, connectorId string, body map[string]interface{}) (err error) {

	_, err = c.SendPostRequest(fmt.Sprintf("/v2/kafka/update-connector/%s", connectorId), body, "Reconfigure Kafka Connector", false)

	return err

}

func deleteConnector(c *client.UpstashClient, connectorId string) (err error) {

	return c.SendDeleteRequest(fmt.Sprintf("/v2/kafka/connector/%s", connectorId), nil, "Delete Kafka Connector", false)

}

func getConnector(c *client.UpstashClient, connectorId string) (connector Connector, err error) {

	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/kafka/connector/%s", connectorId), "Get Kafka Connector", false)

	if err != nil {
		return connector, err
	}

	err = resp.ToJSON(&connector)
	return connector, err
}

func pauseConnector(c *client.UpstashClient, connectorId string) (err error) {

	_, err = c.SendPostRequest(fmt.Sprintf("/v2/kafka/connector/%s/pause", connectorId), nil, "Pause Kafka Connector", false)

	return err

}

func restartConnector(c *client.UpstashClient, connectorId string) (err error) {

	_, err = c.SendPostRequest(fmt.Sprintf("/v2/kafka/connector/%s/restart", connectorId), nil, "Restart Kafka Connector", false)

	return err

}

func startConnector(c *client.UpstashClient, connectorId string) (err error) {

	_, err = c.SendPostRequest(fmt.Sprintf("/v2/kafka/connector/%s/start", connectorId), nil, "Start Kafka Connector", false)

	return err

}
