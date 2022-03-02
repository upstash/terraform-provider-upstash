package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var teamName string
var teamCopyCC bool
var teamMembers map[string]string

func TestUpstashTeamMAIN(t *testing.T) {
	t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey

	teamName = envVars.TeamName
	teamCopyCC = envVars.CopyCC
	teamMembers = envVars.TeamMembers

	terraformOptions := teamOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	teamAsserter(t, terraformOptions)
}

func UpstashTeamRecreate(t *testing.T) {

	teamName = teamName + "Updated"
	teamCopyCC = !teamCopyCC

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

			"teamName":    teamName,
			"copy_cc":     teamCopyCC,
			"teamMembers": teamMembers,
		},
	})

	return terraformOptions
}

func teamAsserter(t *testing.T, terraformOptions *terraform.Options) {
	// fmt.Sprint(map1) == fmt.Sprint(map2)

	nameOutput := terraform.Output(t, terraformOptions, "teamName")
	assert.Equal(t, teamName, nameOutput)

	// Caution here: Copy_cc comes from resource. Not the fetched data.
	ccOutput := terraform.Output(t, terraformOptions, "copy_cc") == "true"
	assert.Equal(t, teamCopyCC, ccOutput)

	// membersOutput := terraform.Output(t, terraformOptions, "teamMembers")
	// assert.Equal(t, createKeyValuePairs(teamMembers), membersOutput)

}

// Obtained from: https://stackoverflow.com/questions/48149969/converting-map-to-string-in-golang
// func createKeyValuePairs(m map[string]string) string {
// 	b := new(bytes.Buffer)
// 	for key, value := range m {
// 		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
// 	}
// 	return b.String()
// }
