
variable "email" {
  description = "Upstash user email"
  default     = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default     = ""
}
provider "upstash" {
  email = var.email
  api_key = var.api_key
}

resource "upstash_qstash_topic" "exampleQstashTopic" {
    name = "terraform_qstash_topic"
}

resource "upstash_qstash_endpoint" "ep" {
    url = "https://testing.com"
    topic_id = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

resource "upstash_qstash_schedule" "sch" {
    destination = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
    cron = "* * * * */2"
    body = '{"key": "value"}'
}

output "topic_id" {
    value = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
}

output "endpoint_id" {
    value = resource.upstash_qstash_endpoint.ep.endpoint_id
}

output "schedule_id" {
    value = resource.upstash_qstash_endpoint.ep.schedule_id
}
