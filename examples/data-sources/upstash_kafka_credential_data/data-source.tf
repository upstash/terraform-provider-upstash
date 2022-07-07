data "upstash_kafka_credential_data" "kafkaCredentialData" {
    credential_id = upstash_kafka_credential.exampleKafkaCredential.credential_id
}
