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

	resp, err := c.SendGetRequest(c.GetQstashEndpoint()+"/schedules/"+scheduleId, "Get QStash Schedule", true)

	if err != nil {
		return schedule, err
	}

	err = resp.ToJSON(&schedule)

	if err != nil {
		return schedule, fmt.Errorf("ERR: %+v\n\n schedule:%+v\n\n resp:%+v", err, schedule, resp)
	}
	return schedule, err
}

func createSchedule(c *client.UpstashClient, body CreateQstashScheduleRequest) (schedule QstashSchedule, err error) {

	err, authorizationToken := c.GetQstashToken()
	if err != nil {
		return schedule, err
	}
	endpoint := c.GetQstashEndpoint() + "/publish/" + body.Destination

	postParameters := []interface{}{
		req.Header{"Upstash-Cron": body.Headers.Cron},
		req.Header{"Upstash-Retries": fmt.Sprint(body.Headers.Retries)},
		req.Header{"Accept": "application/json"},
		req.Header{"Authorization": "Bearer " + authorizationToken},
		req.Header{"Content-Type": body.Headers.ContentType},
		req.Header{"Upstash-Deduplication-Id": body.Headers.DeduplicationId},
		req.Header{"Upstash-Content-Based-Deduplication": fmt.Sprint(body.Headers.ContentBasedDeduplication)},
		req.Header{"Upstash-NotBefore": fmt.Sprint(body.Headers.NotBefore)},
		req.Header{"Upstash-Delay": body.Headers.Delay},
		req.Header{"Upstash-Callback": body.Headers.Callback},
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
		return schedule, err
	}

	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted && resp.Response().StatusCode != http.StatusCreated {
		return schedule, errors.New("Create QStash Schedule failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}

	err = resp.ToJSON(&schedule)
	if err != nil {
		return schedule, fmt.Errorf("ERR: %+v, %+v", resp, err)
	}
	return schedule, err
}

func deleteSchedule(c *client.UpstashClient, scheduleId string) (err error) {
	return c.SendDeleteRequest(c.GetQstashEndpoint()+"/schedules/"+scheduleId, nil, "Delete QStash Schedule", true)
}
