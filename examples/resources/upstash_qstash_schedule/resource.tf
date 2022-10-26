resource "upstash_qstash_schedule" "exampleQstashSchedule" {
    destination = resource.upstash_qstash_topic.exampleQstashTopic.topic_id
    cron = "* * * * */2"

    # or simply provide a link
    # destination = "https://***.***"
}
