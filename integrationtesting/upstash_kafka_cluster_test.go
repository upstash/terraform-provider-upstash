package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var kafkaClusterName, kafkaClusterRegion string
var kafkaClusterMultizone bool

func TestUpstashKafkaClusterMAIN(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey
	kafkaClusterName = envVars.KafkaClusterName
	kafkaClusterRegion = envVars.KafkaClusterRegion
	kafkaClusterMultizone = envVars.KafkaClusterMultiZone

	terraformOptions := kafkaClusterOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaClusterAsserter(t, terraformOptions)

	UpstashKafkaClusterRecreate(t)
	UpstashKafkaClusterUpdate(t)

}

func UpstashKafkaClusterRecreate(t *testing.T) {

	kafkaClusterName = kafkaClusterName + "Updated"
	kafkaClusterRegion = "us-east-1"
	kafkaClusterMultizone = !kafkaClusterMultizone

	terraformOptions := kafkaClusterOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaClusterAsserter(t, terraformOptions)

}

func UpstashKafkaClusterUpdate(t *testing.T) {

	kafkaClusterName = kafkaClusterName + "Updated"

	terraformOptions := kafkaClusterOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaClusterAsserter(t, terraformOptions)
}

func kafkaClusterAsserter(t *testing.T, terraformOptions *terraform.Options) {
	clusterNameOutput := terraform.Output(t, terraformOptions, "cluster_name")
	assert.Equal(t, kafkaClusterName, clusterNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, kafkaClusterRegion, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, kafkaClusterMultizone, multizoneOutput)
}

func kafkaClusterOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_cluster",
		Vars: map[string]interface{}{
			"email":        email,
			"api_key":      apikey,
			"cluster_name": kafkaClusterName,
			"region":       kafkaClusterRegion,
			"multizone":    kafkaClusterMultizone,
		},
	})

	return terraformOptions
}
