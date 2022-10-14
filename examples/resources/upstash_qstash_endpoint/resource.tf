resource "upstash_qstash_endpoint" "exampleQstashEndpoint" {
    url = "https://***.***"
    topic_id = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}