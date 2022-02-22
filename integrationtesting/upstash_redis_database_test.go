package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUpstashRedisDatabase(t *testing.T) {
	t.Parallel()

	upstashCredentials := GetCredentials()

	var (
		email         = upstashCredentials.Email
		apikey        = upstashCredentials.Apikey
		database_name = "TerraformREDISDB"
		region        = "eu-west-1"
		tls           = "true"
		multizone     = "true"
	)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/integrationExamples/upstash_redis_database",
		Vars: map[string]interface{}{
			"email":         email,
			"api_key":       apikey,
			"database_name": database_name,
			"region":        region,
			"tls":           tls,
			"multizone":     multizone,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	// Since using built provider, no need to install from the version
	terraform.Init(t, terraformOptions)

	terraform.Apply(t, terraformOptions)

	terraform.Plan(t, terraformOptions)
}
