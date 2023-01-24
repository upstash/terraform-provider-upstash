
// Topic outputs
output "topic_id" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.topic_id
}
output "topic_name" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.topic_name
}
output "partitions" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.partitions
}

output "retention_time" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.retention_time
}
output "retention_size" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.retention_size
}

output "max_message_size" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.max_message_size
}

output "cleanup_policy" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.cleanup_policy
}

output "creation_time_topic" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.creation_time
}

// cluster outputs
output "cluster_id" {
    value = data.upstash_kafka_cluster_data.kafkaClusterData.cluster_id
}
output "cluster_name" {
    value = data.upstash_kafka_cluster_data.kafkaClusterData.cluster_name
}
output "region" {
    value = data.upstash_kafka_cluster_data.kafkaClusterData.region
}
output "multizone" {
    value = data.upstash_kafka_cluster_data.kafkaClusterData.multizone
}
output "creation_time_cluster" {
    value = data.upstash_kafka_cluster_data.kafkaClusterData.creation_time
}



output "credential_id" {
    value = data.upstash_kafka_credential_data.kafkaCredentialData.credential_id
}

output "credential_name" {
    value = data.upstash_kafka_credential_data.kafkaCredentialData.credential_name
}

output "credential_permissions" {
    value = data.upstash_kafka_credential_data.kafkaCredentialData.permissions
}

output "credential_topic" {
    value = data.upstash_kafka_credential_data.kafkaCredentialData.topic
}

output "credential_username" {
    value = data.upstash_kafka_credential_data.kafkaCredentialData.username
}

output "credential_password" {
    value = data.upstash_kafka_credential_data.kafkaCredentialData.password
}