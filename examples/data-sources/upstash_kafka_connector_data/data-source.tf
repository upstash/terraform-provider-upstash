data "upstash_kafka_connector_data" "kafkaConnectorData" {
  topic_id = resource.upstash_kafka_connector.exampleKafkaConnector.connector_id
}