output "topic_id" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.topic_id
}

output "topic_name" {
    value = data.upstash_kafka_topic_data.kafkaTopicData.topic_name
}

output "cluster_id" {
    value = data.upstash_kafka_cluster_data.kafkaClusterData.cluster_id
}