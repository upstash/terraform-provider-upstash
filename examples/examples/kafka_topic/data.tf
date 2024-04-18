data "upstash_kafka_topic_data" "kafkaTopicData" {
  topic_id = resource.upstash_kafka_topic.exampleKafkaTopic.topic_id
}

data "upstash_kafka_cluster_data" "kafkaClusterData" {
  cluster_id = resource.upstash_kafka_cluster.exampleKafkaCluster.cluster_id
}
