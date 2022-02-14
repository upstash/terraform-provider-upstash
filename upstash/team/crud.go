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
	return resourceRead(ctx, data, m)
}

func resourceRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*client.UpstashClient)
	// teamId := data.Get("team_id").(string)
	// team, err := getTeamMembers(c, teamId)
	team := Team{
		TeamName: "Terraform Team",
		CopyCC:   false,
	}

	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	data.SetId("upstash-team-" + team.TeamId)

	mapping := map[string]interface{}{
		"team_id":   team.TeamId,
		"team_name": team.TeamName,
		"copy_cc":   team.CopyCC,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*client.UpstashClient)
	// teamId := data.Get("team_name").(string)

	// // Only members can change, what to do in that case?
	// if data.HasChange("members") {

	// }

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
