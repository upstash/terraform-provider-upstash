package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var redisDatabaseName, redisDatabaseRegion string
var redisDatabaseTls, redisDatabaseMultizone bool

func TestUpstashRedisDatabaseMAIN(t *testing.T) {
	t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey
	redisDatabaseName = envVars.RedisDatabaseName
	redisDatabaseRegion = envVars.RedisDatabaseRegion
	redisDatabaseTls = envVars.RedisDatabaseTls
	redisDatabaseMultizone = envVars.RedisDatabaseMultiZone

	terraformOptions := redisDatabaseOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

	UpstashRedisDatabaseRecreate(t)
	UpstashRedisDatabaseUpdate(t)

}

func UpstashRedisDatabaseRecreate(t *testing.T) {

	redisDatabaseName = redisDatabaseName + "Updated"
	redisDatabaseRegion = "us-east-1"
	redisDatabaseTls = !redisDatabaseTls
	redisDatabaseMultizone = !redisDatabaseMultizone

	terraformOptions := redisDatabaseOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

}

func UpstashRedisDatabaseUpdate(t *testing.T) {

	redisDatabaseTls = true
	redisDatabaseMultizone = true

	terraformOptions := redisDatabaseOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	redisDatabaseAsserter(t, terraformOptions)

}

func redisDatabaseAsserter(t *testing.T, terraformOptions *terraform.Options) {
	databaseNameOutput := terraform.Output(t, terraformOptions, "database_name")
	assert.Equal(t, redisDatabaseName, databaseNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, redisDatabaseRegion, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, redisDatabaseMultizone, multizoneOutput)
}

func redisDatabaseOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/redis_database",
		Vars: map[string]interface{}{
			"email":         email,
			"api_key":       apikey,
			"database_name": redisDatabaseName,
			"region":        redisDatabaseRegion,
			"tls":           redisDatabaseTls,
			"multizone":     redisDatabaseMultizone,
		},
	})

	return terraformOptions
}
