package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUpstashTeam(t *testing.T) {
	t.Parallel()

	envVars := GetEnvVars()

	var (
		email  = envVars.Email
		apikey = envVars.Apikey

		team_name    = envVars.TeamName
		copy_cc      = envVars.CopyCC
		team_members = envVars.TeamMembers
	)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/team",
		Vars: map[string]interface{}{
			"email":   email,
			"api_key": apikey,

			"team_name":    team_name,
			"copy_cc":      copy_cc,
			"team_members": team_members,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	// Since using built provider, no need to install from the version
	// terraform.Init(t, terraformOptions)

	terraform.Apply(t, terraformOptions)

	terraform.Plan(t, terraformOptions)
}
