package team

import (
	"github.com/imroc/req"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
)

func createTeam(c *client.UpstashClient, body CreateTeamRequest) (team Team, err error) {

	resp, err := c.SendPostRequest("/v2/team", body, "Create Team")

	if err != nil {
		return team, err
	}

	err = resp.ToJSON(&team)
	return team, err

}

func getTeamMembers(c *client.UpstashClient, teamId string) (teamMembers []GetTeamMembers, err error) {

	resp, err := c.SendGetRequest("/v2/teams/"+teamId, "Get Team Members")

	if err != nil {
		return teamMembers, err
	}

	err = resp.ToJSON(&teamMembers)
	return teamMembers, err

}

func deleteTeam(c *client.UpstashClient, teamId string) (err error) {

	return c.SendDeleteRequest("/v2/team/"+teamId, nil, "Delete Team")

}

func addMember(c *client.UpstashClient, teamId string, email string, role string) (err error) {

	param := req.Param{
		"team_id":      teamId,
		"member_email": email,
		"member_role":  role,
	}

	_, err = c.SendPostRequest("/v2/teams/member", param, "Add Member to Team")

	return err

}

func removeMember(c *client.UpstashClient, teamId string, email string) (err error) {

	param := req.Param{
		"team_id":      teamId,
		"member_email": email,
	}

	return c.SendDeleteRequest("/v2/teams/member", param, "Remove Team Member")

}
