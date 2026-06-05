package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var redis_database_name, redis_database_platform string

func expectedRegionForPlatform(platform string) string {
	switch platform {
	case "gcp":
		return "gcp-global"
	case "aws":
		return "global"
	default:
		return ""
	}
}

func TestUpstashRedisDatabaseMAIN(t *testing.T) {
	t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey
	redis_database_name = envVars.RedisDatabaseName
	redis_database_platform = envVars.RedisDatabasePlatform
	if redis_database_platform == "" {
		redis_database_platform = "aws"
	}

	terraformOptions := redisDatabaseOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

	UpstashRedisDatabaseRecreate(t)
	UpstashRedisDatabaseUpdate(t)

}

func UpstashRedisDatabaseRecreate(t *testing.T) {

	// Changing the (ForceNew) database_name triggers recreation of the resource.
	redis_database_name = redis_database_name + "Updated"

	terraformOptions := redisDatabaseOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

}

func UpstashRedisDatabaseUpdate(t *testing.T) {

	// Re-apply with the same configuration to confirm the resource is stable
	// (no perpetual diff for the computed platform/region fields).
	terraformOptions := redisDatabaseOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

}

func redisDatabaseAsserter(t *testing.T, terraformOptions *terraform.Options) {
	databaseNameOutput := terraform.Output(t, terraformOptions, "database_name")
	assert.Equal(t, redis_database_name, databaseNameOutput)

	platformOutput := terraform.Output(t, terraformOptions, "platform")
	assert.Equal(t, redis_database_platform, platformOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, expectedRegionForPlatform(redis_database_platform), regionOutput)
}

func redisDatabaseOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/redis_database",
		Vars: map[string]interface{}{
			"email":         email,
			"api_key":       apikey,
			"database_name": redis_database_name,
			"platform":      redis_database_platform,
		},
	})

	return terraformOptions
}
