resource "upstash_qstash_topic_v2" "exampleQstashTopicNew" {
    name = "terraform_qstash_topic"
    endpoints = ["https://google.com"]
}

resource "upstash_qstash_schedule_v2" "exampleQstashSchedule" {
    destination = upstash_qstash_topic_v2.exampleQstashTopicNew.name
    cron = "* * * * */2"
    delay = 3600
}

resource "upstash_qstash_schedule_v2" "exampleQstashSchedule2" {
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

data "upstash_qstash_schedule_v2_data" "qstashData" {
    schedule_id = upstash_qstash_schedule_v2.exampleQstashSchedule2.schedule_id
}

output exampleOutput {
    value = data.upstash_qstash_schedule_v2_data.qstashData
}
