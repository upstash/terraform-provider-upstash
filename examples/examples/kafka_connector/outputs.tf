output "state" {
    value=data.upstash_kafka_connector_data.exampleKafkaConnectorData.state
}

output "connector_state" {
    value=data.upstash_kafka_connector_data.exampleKafkaConnectorData.connector_state
}