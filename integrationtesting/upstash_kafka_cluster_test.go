package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var email, apikey, cluster_name, region string
var multizone bool

func TestUpstashKafkaClusterMAIN(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey
	cluster_name = envVars.KafkaClusterName
	region = envVars.KafkaClusterRegion
	multizone = envVars.KafkaClusterMultiZone

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_cluster",
		Vars: map[string]interface{}{
			"email":        email,
			"api_key":      apikey,
			"cluster_name": cluster_name,
			"region":       region,
			"multizone":    multizone,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	// Since using built provider, no need to install from the version
	// terraform.Init(t, terraformOptions)

	terraform.Apply(t, terraformOptions)

	terraform.Plan(t, terraformOptions)

	clusterNameOutput := terraform.Output(t, terraformOptions, "cluster_name")
	assert.Equal(t, cluster_name, clusterNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, region, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, multizone, multizoneOutput)

	UpstashKafkaClusterRecreate(t)
	UpstashKafkaClusterUpdate(t)

}

func UpstashKafkaClusterRecreate(t *testing.T) {

	cluster_name = cluster_name + "Updated"
	region = "us-east-1"
	multizone = !multizone

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_cluster",
		Vars: map[string]interface{}{
			"email":        email,
			"api_key":      apikey,
			"cluster_name": cluster_name,
			"region":       region,
			"multizone":    multizone,
		},
	})

	terraform.Apply(t, terraformOptions)

	terraform.Plan(t, terraformOptions)

	clusterNameOutput := terraform.Output(t, terraformOptions, "cluster_name")
	assert.Equal(t, cluster_name, clusterNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, region, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, multizone, multizoneOutput)

}

func UpstashKafkaClusterUpdate(t *testing.T) {
	cluster_name = cluster_name + "Updated"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_cluster",
		Vars: map[string]interface{}{
			"email":        email,
			"api_key":      apikey,
			"cluster_name": cluster_name,
			"region":       region,
			"multizone":    multizone,
		},
	})

	terraform.Apply(t, terraformOptions)

	terraform.Plan(t, terraformOptions)

	clusterNameOutput := terraform.Output(t, terraformOptions, "cluster_name")
	assert.Equal(t, cluster_name, clusterNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, region, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, multizone, multizoneOutput)
}
