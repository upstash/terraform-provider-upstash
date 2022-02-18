package team

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	team, err := createTeam(c, CreateTeamRequest{
		TeamName: data.Get("team_name").(string),
		CopyCC:   data.Get("copy_cc").(bool),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-team-" + team.TeamId)
	data.Set("team_id", team.TeamId)

	return resourceUpdate(ctx, data, m)
	// return resourceRead(ctx, data, m)
}

func resourceRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	teamId := data.Get("team_id").(string)
	teamMembers, err := getTeamMembers(c, teamId)

	teamName := teamMembers[0].TeamName

	membersMap := make(map[string]string)

	for i := 0; i < len(teamMembers); i++ {
		membersMap[teamMembers[i].MemberEmail] = teamMembers[i].MemberRole
	}

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-team-" + teamId)

	mapping := map[string]interface{}{
		"team_id":      teamId,
		"team_name":    teamName,
		"team_members": membersMap,
		// "copy_cc": data.Get("copy_cc").(bool)
	}

	return utils.SetAndCheckErrors(data, mapping)

}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	teamId := data.Get("team_id").(string)

	// Only members can change, what to do in that case?
	if data.HasChange("team_members") {
		membersMap := data.Get("team_members").(map[string]interface{})

		for email, role := range membersMap {
			if role != "owner" {

				memberNotFound, err := removeMember(c, teamId, email)
				if err != nil && !memberNotFound {
					return diag.FromErr(err)
				}

				err = addMember(c, teamId, email, role.(string))
				if err != nil {
					return diag.FromErr(err)
				}

			}
		}

	}

	// TO DO: ADD MEMBER ADD/REMOVE LOGIC HERE.

	return resourceRead(ctx, data, m)
}

func resourceDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	teamId := data.Get("team_id").(string)
	err := deleteTeam(c, teamId)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
