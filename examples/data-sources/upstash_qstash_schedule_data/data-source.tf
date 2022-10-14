data "upstash_qstash_schedule_data" "exampleQstashScheduleData" {
    schedule_id = resource.upstash_qstash_schedule.exampleQstashSchedule.schedule_id
}