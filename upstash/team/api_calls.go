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

func createTeam(c *client.UpstashClient, body CreateTeamRequest) (team Team, err error) {
	resp, err := req.Post(api_endpoint+"/v2/team",
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
		req.BodyJSON(body),
	)
	if err != nil {
		return team, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return team, errors.New("Create team failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&team)
	return team, err
}

func getTeamMembers(c *client.UpstashClient, teamId string) (team Team, err error) {
	resp, err := req.Get(api_endpoint+"/v2/teams/"+teamId,
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

func deleteTeam(c *client.UpstashClient, clusterId string) (err error) {
	resp, err := req.Delete(api_endpoint+"/v2/team/"+clusterId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New("Delete team failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())

	}
	return err
}
