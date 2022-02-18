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

func getTeamMembers(c *client.UpstashClient, teamId string) (teamMembers []GetTeamMembers, err error) {
	resp, err := req.Get(api_endpoint+"/v2/teams/"+teamId,
		req.Header{
			"Accept":        "application/json",
			"Authorization": utils.BasicAuth(c.Email, c.Apikey),
		},
	)
	if err != nil {
		return teamMembers, err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return teamMembers, errors.New("Get team members failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}
	err = resp.ToJSON(&teamMembers)
	return teamMembers, err
}

func deleteTeam(c *client.UpstashClient, teamId string) (err error) {
	resp, err := req.Delete(api_endpoint+"/v2/team/"+teamId,
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

func addMember(c *client.UpstashClient, teamId string, email string, role string) (err error) {

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": utils.BasicAuth(c.Email, c.Apikey),
	}

	param := req.Param{
		"team_id":      teamId,
		"member_email": email,
		"member_role":  role,
	}

	resp, err := req.Post(api_endpoint+"/v2/teams/member", req.BodyJSON(param), header)

	if err != nil {
		return err
	}
	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {
		return errors.New(teamId + " " + email + " " + role + " " + "Add team member failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	}

	// return errors.New(teamId + " " + email + " " + role + " " + "Add team member succeeded, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())
	return nil

}

func removeMember(c *client.UpstashClient, teamId string, email string) (memberNotFound bool, err error) {

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": utils.BasicAuth(c.Email, c.Apikey),
	}

	param := req.Param{
		"team_id":      teamId,
		"member_email": email,
	}

	resp, err := req.Delete(api_endpoint+"/v2/teams/member", req.BodyJSON(param), header)

	if err != nil {
		return false, err
	}

	if resp.Response().StatusCode != http.StatusOK && resp.Response().StatusCode != http.StatusAccepted {

		// Instead of this, use couldn't find
		if resp.String() == `"Couldn't find the member"` {
			return true, nil
		}

		return false, errors.New(teamId + " " + email + " " + "Remove team member failed, status code: " + strconv.Itoa(resp.Response().StatusCode) + " response: " + resp.String())

	}
	return false, nil

}
