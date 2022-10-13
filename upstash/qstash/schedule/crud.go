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

	mapping := map[string]interface{}{
		"content":     schedule.Content,
		"createdAt":   schedule.CreatedAt,
		"cron":        schedule.Cron,
		"destination": schedule.Destination,
		"scheduleId":  schedule.ScheduleId,
		"settings":    schedule.Settings,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceScheduleCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	schedule, err := createSchedule(c, CreateQstashScheduleRequest{
		Destination:      data.Get("destination").(string),
		HeaderParameters: data.Get("headerParameters").(HeaderParameters),
	})
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
