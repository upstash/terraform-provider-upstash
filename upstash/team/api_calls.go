package team

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

const api_endpoint = client.UPSTASH_API_ENDPOINT

func getTeamMembers(c *client.UpstashClient, teamId string) (team Team, err error) {
	resp, err := req.Get(api_endpoint+"/v2/team/"+teamId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return team, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return team, errors.New("Get team members failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&team)
	return team, err
}
