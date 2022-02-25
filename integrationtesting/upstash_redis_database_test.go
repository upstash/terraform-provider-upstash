package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// var email, apikey, database_name, region string
// var tls, multizone bool

func TestUpstashRedisDatabaseMAIN(t *testing.T) {
	t.Parallel()

	envVars := GetEnvVars()

	var (
		email         = envVars.Email
		apikey        = envVars.Apikey
		database_name = envVars.RedisDatabaseName
		region        = envVars.RedisDatabaseRegion
		tls           = envVars.RedisDatabaseTls
		multizone     = envVars.RedisDatabaseMultiZone
	)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/redis_database",
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
	// terraform.Init(t, terraformOptions)

	terraform.Apply(t, terraformOptions)

	terraform.Plan(t, terraformOptions)

}
