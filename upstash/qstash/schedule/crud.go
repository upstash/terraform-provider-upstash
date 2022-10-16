package schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceScheduleRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	scheduleId := data.Get("schedule_id").(string)
	schedule, err := getSchedule(c, scheduleId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-schedule-" + schedule.ScheduleId)

	destination := schedule.Destination.Topic.TopicId
	if schedule.Destination.Type == "url" {
		destination = schedule.Destination.Url
	}
	mapping := map[string]interface{}{
		"created_at":  schedule.CreatedAt,
		"retries":     schedule.Settings.Retries,
		"not_before":  schedule.Settings.NotBefore,
		"cron":        schedule.Cron,
		"destination": destination,
		"schedule_id": schedule.ScheduleId,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceScheduleCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)

	schedule, err := createSchedule(c, CreateQstashScheduleRequest{
		Destination: data.Get("destination").(string),
		Body:        data.Get("body").(string),
	},
		data.Get("content_type").(string),
		data.Get("deduplication_id").(string),
		data.Get("content_based_deduplication").(bool),
		data.Get("not_before").(int),
		data.Get("delay").(string),
		data.Get("retries").(int),
		data.Get("cron").(string),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-schedule-" + schedule.ScheduleId)
	data.Set("schedule_id", schedule.ScheduleId)

	return resourceScheduleRead(ctx, data, m)
}

func resourceScheduleDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	scheduleId := data.Get("schedule_id").(string)
	err := deleteSchedule(c, scheduleId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
