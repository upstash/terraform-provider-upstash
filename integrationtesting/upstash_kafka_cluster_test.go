package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var kafka_cluster_name, kafka_cluster_region string
var kafka_cluster_multizone bool

func TestUpstashKafkaClusterMAIN(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey
	kafka_cluster_name = envVars.KafkaClusterName
	kafka_cluster_region = envVars.KafkaClusterRegion
	kafka_cluster_multizone = envVars.KafkaClusterMultiZone

	terraformOptions := kafkaClusterOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaClusterAsserter(t, terraformOptions)

	UpstashKafkaClusterRecreate(t)
	UpstashKafkaClusterUpdate(t)

}

func UpstashKafkaClusterRecreate(t *testing.T) {

	kafka_cluster_name = kafka_cluster_name + "Updated"
	kafka_cluster_region = "us-east-1"
	kafka_cluster_multizone = !kafka_cluster_multizone

	terraformOptions := kafkaClusterOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaClusterAsserter(t, terraformOptions)

}

func UpstashKafkaClusterUpdate(t *testing.T) {

	kafka_cluster_name = kafka_cluster_name + "Updated"

	terraformOptions := kafkaClusterOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaClusterAsserter(t, terraformOptions)
}

func kafkaClusterAsserter(t *testing.T, terraformOptions *terraform.Options) {
	clusterNameOutput := terraform.Output(t, terraformOptions, "cluster_name")
	assert.Equal(t, kafka_cluster_name, clusterNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, kafka_cluster_region, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, kafka_cluster_multizone, multizoneOutput)
}

func kafkaClusterOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_cluster",
		Vars: map[string]interface{}{
			"email":        email,
			"api_key":      apikey,
			"cluster_name": kafka_cluster_name,
			"region":       kafka_cluster_region,
			"multizone":    kafka_cluster_multizone,
		},
	})

	return terraformOptions
}
