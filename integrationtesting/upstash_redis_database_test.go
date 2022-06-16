package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var redis_database_name, redis_database_region string
var redis_database_tls, redis_database_multizone bool

func TestUpstashRedisDatabaseMAIN(t *testing.T) {
	t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey
	redis_database_name = envVars.RedisDatabaseName
	redis_database_region = envVars.RedisDatabaseRegion
	redis_database_tls = envVars.RedisDatabaseTls
	redis_database_multizone = envVars.RedisDatabaseMultiZone

	terraformOptions := redisDatabaseOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

	UpstashRedisDatabaseRecreate(t)
	UpstashRedisDatabaseUpdate(t)

}

func UpstashRedisDatabaseRecreate(t *testing.T) {

	redis_database_name = redis_database_name + "Updated"
	redis_database_region = "us-east-1"
	redis_database_tls = !redis_database_tls
	redis_database_multizone = !redis_database_multizone

	terraformOptions := redisDatabaseOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

}

func UpstashRedisDatabaseUpdate(t *testing.T) {

	redis_database_tls = true
	redis_database_multizone = true

	terraformOptions := redisDatabaseOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

}

func redisDatabaseAsserter(t *testing.T, terraformOptions *terraform.Options) {
	databaseNameOutput := terraform.Output(t, terraformOptions, "database_name")
	assert.Equal(t, redis_database_name, databaseNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, redis_database_region, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, redis_database_multizone, multizoneOutput)
}

func redisDatabaseOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/redis_database",
		Vars: map[string]interface{}{
			"email":         email,
			"api_key":       apikey,
			"database_name": redis_database_name,
			"region":        redis_database_region,
			"multizone":     redis_database_multizone,
		},
	})

	return terraformOptions
}
