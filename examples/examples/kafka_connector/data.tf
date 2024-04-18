data "upstash_kafka_connector_data" "exampleKafkaConnectorData" {
  connector_id = upstash_kafka_connector.exampleKafkaConnector.connector_id
}

