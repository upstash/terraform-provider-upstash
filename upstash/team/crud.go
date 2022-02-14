package team

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	teamId := data.Get("team_id").(string)
	team, err := getTeamMembers(c, teamId)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-team-" + team.TeamId)

	mapping := map[string]interface{}{
		"team_id":   team.TeamId,
		"team_name": team.TeamName,
		"copy_cc":   team.CopyCC,
	}

	return utils.SetAndCheckErrors(data, mapping)
}
