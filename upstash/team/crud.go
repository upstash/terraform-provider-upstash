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

	var copycc bool
	for i := 0; i < len(teamMembers); i++ {
		membersMap[teamMembers[i].MemberEmail] = teamMembers[i].MemberRole
		if teamMembers[i].CopyCC {
			copycc = teamMembers[i].CopyCC
		}
	}

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-team-" + teamId)

	mapping := map[string]interface{}{
		"team_id":      teamId,
		"team_name":    teamName,
		"team_members": membersMap,
		"copy_cc":      copycc,
	}

	return utils.SetAndCheckErrors(data, mapping)

}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	teamId := data.Get("team_id").(string)

	// Only members can change, what to do in that case?
	if data.HasChange("team_members") {

		a, b := data.GetChange("team_members")
		old := a.(map[string]interface{})
		new := b.(map[string]interface{})

		for email, role := range old {
			if role.(string) != "owner" {
				if role.(string) != new[email] {
					err := removeMember(c, teamId, email)
					if err != nil {
						return diag.FromErr(err)
					}
				}
			}
		}

		for email, role := range new {
			if role.(string) != "owner" {
				if role.(string) != old[email] {
					err := addMember(c, teamId, email, role.(string))
					if err != nil {
						return diag.FromErr(err)
					}
				}
			}
		}
	}

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
