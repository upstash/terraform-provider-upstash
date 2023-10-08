resource "upstash_qstash_topic" "exampleQstashTopicNew" {
    name = "terraform_qstash_topic"
    endpoints = ["https://google.com"]
}

resource "upstash_qstash_schedule" "exampleQstashSchedule" {
    destination = upstash_qstash_topic.exampleQstashTopicNew.name
    cron = "* * * * */2"
    delay = 3600
}

resource "upstash_qstash_schedule" "exampleQstashSchedule2" {
    destination = "https://google.com"
    cron = "* * * * */3"
    forward_headers = {
        My-Header : "My-value"
        My-Header2 : "My-value2"
    }
    callback = "https://testing-url.com"
    body = "{\"key\": \"value\"}"
    method = "GET"
}

data "upstash_qstash_schedule_data" "kafkaClusterData" {
    schedule_id = upstash_qstash_schedule.exampleQstashSchedule2.schedule_id
}

output a {
    value = data.upstash_qstash_schedule_data.kafkaClusterData
}
