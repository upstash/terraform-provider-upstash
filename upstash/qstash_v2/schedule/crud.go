package schedule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceScheduleRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	scheduleId := data.Get("schedule_id").(string)
	if scheduleId == "" {
		scheduleId = data.Id()
	}

	schedule, err := getSchedule(c, scheduleId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-schedule-" + schedule.ScheduleId)

	mapping := map[string]interface{}{
		"created_at":  schedule.CreatedAt,
		"schedule_id": schedule.ScheduleId,
		"cron":        schedule.Cron,
		"destination": schedule.Destination,
		"method":      schedule.Method,
		"header":      fmt.Sprintf("%+v", schedule.Header),
		"body":        schedule.Body,
		"retries":     schedule.Retries,
		"delay":       fmt.Sprintf("%d", schedule.Delay),
		"callback":    schedule.Callback,
	}

	return utils.SetAndCheckErrors(data, mapping)
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

func resourceScheduleCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)

	scheduleID, err := createSchedule(c, CreateQstashScheduleRequest{
		Destination:    data.Get("destination").(string),
		Body:           data.Get("body").(string),
		ForwardHeaders: data.Get("forward_headers").(map[string]interface{}),
		Headers: QstashScheduleHeaders{
			ContentType: data.Get("content_type").(string),
			Method:      data.Get("method").(string),
			Delay:       data.Get("delay").(string) + "s",
			Retries:     data.Get("retries").(int),
			Callback:    data.Get("callback").(string),
			Cron:        data.Get("cron").(string),
		},
	})
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-qstash-schedule-" + scheduleID)
	data.Set("schedule_id", scheduleID)

	return resourceScheduleRead(ctx, data, m)
}
