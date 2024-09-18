data "upstash_kafka_topic_data" "kafkaTopicData" {
  topic_id = resource.upstash_kafka_topic.exampleKafkaTopic.topic_id
}