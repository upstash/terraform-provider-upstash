
resource "upstash_qstash_topic" "exampleQstashTopic" {
    name = "terraform_qstash_topic_rename"
}

output "a" {
    value = "TESTTTTT"
}

output "b" {
    value = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

resource "upstash_qstash_endpoint" "ep" {
    url = "https://google.com"
    topic_id = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

resource "upstash_qstash_endpoint" "ep2" {
    url = "https://testing.com"
    topic_id = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

resource "upstash_qstash_schedule" "sch1" {
    destination = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
    #for now, runs default * * * * *
    cron = "* * * * * invalid"
}

resource "upstash_qstash_schedule" "sch2" {
    destination = "https://testing.com"
    #for now, runs default * * * * *
    cron = "* * * * * invalid"
}

output "a2" {
    value = "TESTTTTT"
}

output "b2" {
    value = resource.upstash_qstash_endpoint.ep.endpoint_id
}

output "b3" {
    value = resource.upstash_qstash_endpoint.ep2.endpoint_id
}