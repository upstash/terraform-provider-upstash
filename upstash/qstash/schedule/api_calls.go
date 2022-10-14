package schedule

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func getSchedule(c *client.UpstashClient, scheduleId string) (schedule QstashSchedule, err error) {

	resp, _ := c.SendGetRequest(c.GetQstashEndpoint()+"/schedules/"+scheduleId, "Get QStash Schedule")

	if err != nil {
		return schedule, err
	}

	err = resp.ToJSON(&schedule)
	if err != nil {
		return schedule, fmt.Errorf("ERR: %+v\n\n schedule:%+v\n\n resp:%+v", err, schedule, resp)
	}
	return schedule, err
}

func createSchedule(c *client.UpstashClient, body CreateQstashScheduleRequest, contentType string, deduplicationId string, contentBasedDeduplication bool, notBefore int, delay string, retries int, cron string) (schedule QstashSchedule, err error) {

	err, BEARER_TOKEN := c.GetQstashToken()
	if err != nil {
		return schedule, err
	}
	endpoint := c.GetQstashEndpoint() + "/publish/" + body.Destination

	resp, err := req.Post(
		endpoint,
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": "Bearer " + BEARER_TOKEN},
		req.Header{"Content-Type": contentType},
		req.Header{"Upstash-Deduplication-Id": deduplicationId},
		req.Header{"Upstash-Content-Based-Deduplication": fmt.Sprint(contentBasedDeduplication)},
		req.Header{"Upstash-NotBefore": fmt.Sprint(notBefore)},
		req.Header{"Upstash-Delay": delay},
		req.Header{"Upstash-Retries": fmt.Sprint(retries)},
		req.Header{"Upstash-Cron": cron},
		req.BodyJSON(body.Body),
	)

	if err != nil {
		return schedule, err
	}

	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return schedule, errors.New("Create QStash Schedule failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}

	err = resp.ToJSON(&schedule)
	if err != nil {
		return schedule, fmt.Errorf("ERR2: %+v, %+v", resp, err)
	}
	return schedule, err
}

func deleteSchedule(c *client.UpstashClient, scheduleId string) (err error) {
	return c.SendDeleteRequest(c.GetQstashEndpoint()+"/schedules/"+scheduleId, nil, "Delete QStash Schedule")
}
