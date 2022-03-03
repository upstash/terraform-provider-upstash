package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var team_name string
var team_copy_cc bool
var team_members map[string]string

func TestUpstashTeamMAIN(t *testing.T) {
	t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey

	team_name = envVars.TeamName
	team_copy_cc = envVars.CopyCC
	team_members = envVars.TeamMembers

	terraformOptions := teamOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	teamAsserter(t, terraformOptions)
}

func UpstashTeamRecreate(t *testing.T) {

	team_name = team_name + "Updated"
	team_copy_cc = !team_copy_cc

	terraformOptions := teamOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	teamAsserter(t, terraformOptions)

}

func teamOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/team",
		Vars: map[string]interface{}{
			"email":   email,
			"api_key": apikey,

			"team_name":    team_name,
			"copy_cc":      team_copy_cc,
			"team_members": team_members,
		},
	})

	return terraformOptions
}

func teamAsserter(t *testing.T, terraformOptions *terraform.Options) {
	// fmt.Sprint(map1) == fmt.Sprint(map2)

	nameOutput := terraform.Output(t, terraformOptions, "team_name")
	assert.Equal(t, team_name, nameOutput)

	// Caution here: Copy_cc comes from resource. Not the fetched data.
	ccOutput := terraform.Output(t, terraformOptions, "copy_cc") == "true"
	assert.Equal(t, team_copy_cc, ccOutput)

	// membersOutput := terraform.Output(t, terraformOptions, "team_members")
	// assert.Equal(t, createKeyValuePairs(team_members), membersOutput)

}

// Obtained from: https://stackoverflow.com/questions/48149969/converting-map-to-string-in-golang
// func createKeyValuePairs(m map[string]string) string {
// 	b := new(bytes.Buffer)
// 	for key, value := range m {
// 		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
// 	}
// 	return b.String()
// }
