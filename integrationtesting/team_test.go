package integrationtesting

// import (
// 	"testing"

// 	"github.com/gruntwork-io/terratest/modules/terraform"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUpstashTeam(t *testing.T) {
// 	t.Parallel()

// 	terraformOptions := &terraform.Options{
// 		TerraformDir: "../test/team",
// 	}

// 	defer terraform.Destroy(t, terraformOptions)

// 	// terraform.InitAndApply(t, terraformOptions)
// 	terraform.Plan(t, terraformOptions)
// 	terraform.Apply(t, terraformOptions)

// 	teamName := terraform.Output(t, terraformOptions, "team_name")

// 	assert.Equal(t, "Terraform Team I1", teamName)
// }
