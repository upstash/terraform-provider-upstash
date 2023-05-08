package connector

import (
	"fmt"

	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func createConnector(c *client.UpstashClient, clusterId string, body CreateConnectorRequest) (connector Connector, err error) {

	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/kafka/cluster/%s/connector", clusterId), body, "Create Kafka Connector", false)

	if err != nil {
		return connector, err
	}

	err = resp.ToJSON(&connector)
	return connector, err

}

func reconfigureConnector(c *client.UpstashClient, clusterId, connectorId string, body map[string]interface{}) (err error) {

	_, err = c.SendPatchRequest(fmt.Sprintf("/v2/kafka/cluster/%s/connector/%s", clusterId, connectorId), body, "Reconfigure Kafka Cluster", false)

	return err

}

func deleteConnector(c *client.UpstashClient, clusterId, connectorId string) (err error) {

	return c.SendDeleteRequest(fmt.Sprintf("/v2/kafka/cluster/%s/connector/%s", clusterId, connectorId), nil, "Delete Kafka Connector", false)

}

func getConnector(c *client.UpstashClient, connectorId string) (connector Connector, err error) {

	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/kafka/connector/%s", connectorId), "Get Kafka Connector", false)

	if err != nil {
		return connector, err
	}

	err = resp.ToJSON(&connector)
	return connector, err
}

func pauseConnector(c *client.UpstashClient, clusterId, connectorId string) (err error) {

	_, err = c.SendPostRequest(fmt.Sprintf("/v2/kafka/cluster/%s/connector/%s/pause", clusterId, connectorId), nil, "Pause Kafka Connector", false)

	return err

}

func restartConnector(c *client.UpstashClient, clusterId, connectorId string) (err error) {

	_, err = c.SendPostRequest(fmt.Sprintf("/v2/kafka/cluster/%s/connector/%s/restart", clusterId, connectorId), nil, "Restart Kafka Connector", false)

	return err

}

func startConnector(c *client.UpstashClient, clusterId, connectorId string) (err error) {

	_, err = c.SendPostRequest(fmt.Sprintf("/v2/kafka/cluster/%s/connector/%s/start", clusterId, connectorId), nil, "Start Kafka Connector", false)

	return err

}
