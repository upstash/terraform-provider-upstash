resource "upstash_qstash_topic" "exampleQstashTopic" {
    name = "terraform_qstash_topic"
}

resource "upstash_qstash_endpoint" "ep" {
    url = "https://testing.com"
    topic_id = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

resource "upstash_qstash_endpoint" "ep2" {
    url = "https://testing2.com"
    topic_id = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

resource "upstash_qstash_schedule" "sch" {
    destination = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
    cron = "* * * * */2"
    body = "{\"key\": \"value\"}"
    forward_headers = {
        My-Header : "My-value"
        My-Header2 : "My-value2"
    }
}

output "topic_id" {
    value = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

output "endpoint_id" {
    value = resource.upstash_qstash_endpoint.ep.endpoint_id
}

output "schedule_id" {
    value = resource.upstash_qstash_schedule.sch.schedule_id
}

output "endpoints_of_topic" {
    value = resource.upstash_qstash_topic.exampleQstashTopic.endpoints
}

output "endpoints_of_topic2" {
    value = resource.upstash_qstash_schedule.sch.body
}