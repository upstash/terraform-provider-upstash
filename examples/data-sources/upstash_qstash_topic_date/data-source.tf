data "upstash_qstash_topic_data" "exampleQstashTopicData" {
  topic_id = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}