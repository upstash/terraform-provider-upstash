package schedule

import (
	"fmt"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

var QSTASH_API_ENDPOINT = "https://qstash.upstash.io/v1"

func getSchedule(c *client.UpstashClient, scheduleId string) (schedule QstashSchedule, err error) {

	resp, _ := c.SendGetRequest(QSTASH_API_ENDPOINT+"/schedules/"+scheduleId, "Get QStash Schedule")

	if err != nil {
		return schedule, err
	}

	err = resp.ToJSON(&schedule)
	if err != nil {
		return schedule, fmt.Errorf("ERR: %+v, %+v", resp, err)
	}
	return schedule, err
}

func createSchedule(c *client.UpstashClient, body CreateQstashScheduleRequest, headers ...req.Header) (schedule QstashSchedule, err error) {
	resp, _ := c.SendPostRequest(QSTASH_API_ENDPOINT+"/publish/"+body.Destination, body, "Create QStash Schedule", headers...)

	if err != nil {
		return schedule, err
	}

	err = resp.ToJSON(&schedule)
	if err != nil {
		return schedule, fmt.Errorf("ERR2: %+v, %+v", resp, err)
	}
	return schedule, err
}

func deleteSchedule(c *client.UpstashClient, scheduleId string) (err error) {
	return c.SendDeleteRequest(QSTASH_API_ENDPOINT+"/schedules/"+scheduleId, nil, "Delete QStash Schedule")
}
