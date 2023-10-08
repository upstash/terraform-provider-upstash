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

	resp, _ := c.SendGetRequest(c.GetQstashEndpoint()+"/schedules/"+scheduleId, "Get QStash Schedule", true)

	if err != nil {
		return schedule, err
	}

	err = resp.ToJSON(&schedule)
	if err != nil {
		return schedule, fmt.Errorf("ERR: %+v\n\n schedule:%+v\n\n resp:%+v", err, schedule, resp)
	}
	return schedule, err
}

func deleteSchedule(c *client.UpstashClient, scheduleId string) (err error) {
	return c.SendDeleteRequest(c.GetQstashEndpoint()+"/schedules/"+scheduleId, nil, "Delete QStash Schedule", true)
}

func createSchedule(c *client.UpstashClient, body CreateQstashScheduleRequest) (scheduleID string, err error) {

	err, BEARER_TOKEN := c.GetQstashToken()
	if err != nil {
		return "", err
	}
	endpoint := c.GetQstashEndpoint() + "/schedules/" + body.Destination

	postParameters := []interface{}{
		req.Header{"Content-Type": body.Headers.ContentType},
		req.Header{"Upstash-Method": body.Headers.Method},
		req.Header{"Upstash-Delay": body.Headers.Delay},
		req.Header{"Upstash-Retries": fmt.Sprint(body.Headers.Retries)},
		req.Header{"Upstash-Callback": body.Headers.Callback},
		req.Header{"Upstash-Cron": body.Headers.Cron},

		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": "Bearer " + BEARER_TOKEN},

		req.BodyJSON(body.Body),
	}

	forwardHeaders := body.ForwardHeaders
	for index := range forwardHeaders {
		postParameters = append(postParameters, req.Header{fmt.Sprintf("Upstash-Forward-%s", index): forwardHeaders[index].(string)})
	}

	resp, err := req.Post(
		endpoint,
		postParameters...,
	)

	if err != nil {
		return "", err
	}

	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return "", errors.New("Create QStash Schedule failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}

	var response CreateQstashScheduleResponse
	err = resp.ToJSON(&response)
	if err != nil {
		return scheduleID, fmt.Errorf("ERR: %+v, %+v", resp, err)
	}

	return response.ScheduleId, err
}
