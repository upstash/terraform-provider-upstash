package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestUpstashTeam(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../test/team",
	}

	defer terraform.Destroy(t, terraformOptions)

	// terraform.InitAndApply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

	teamName := terraform.Output(t, terraformOptions, "cluster_name")

	assert.Equal(t, "terraformish Teams", teamName)
}
