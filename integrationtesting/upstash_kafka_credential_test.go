package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var kafka_credential_topic_name, kafka_credential_topic_cleanup_policy string
var kafka_credential_topic_partitions, kafka_credential_topic_retention_time, kafka_credential_topic_retention_size, kafka_credential_topic_max_message_size int

var credential_cluster_name, credential_cluster_region string
var credential_cluster_multizone bool

var credential_name, credential_permissions string

func TestUpstashKafkaCredentialMAIN(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey

	credential_cluster_name = envVars.KafkaClusterName
	credential_cluster_region = envVars.KafkaClusterRegion
	credential_cluster_multizone = envVars.KafkaClusterMultiZone

	kafka_credential_topic_name = envVars.KafkaTopicName
	kafka_credential_topic_partitions = envVars.KafkaTopicPartitions
	kafka_credential_topic_retention_time = envVars.KafkaTopicRetentionTime
	kafka_credential_topic_retention_size = envVars.KafkaTopicRetentionSize
	kafka_credential_topic_max_message_size = envVars.KafkaTopicMaxMessageSize
	kafka_credential_topic_cleanup_policy = envVars.KafkaTopicCleanupPolicy

	credential_name = envVars.KafkaCredentialName
	credential_permissions = envVars.KafkaCredentialPermissions

	terraformOptions := kafkaCredentialOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaCredentialAsserter(t, terraformOptions)
}

func kafkaCredentialAsserter(t *testing.T, terraformOptions *terraform.Options) {

	credentialPermissionsOutput := terraform.Output(t, terraformOptions, "credential_permissions")
	assert.Equal(t, credential_permissions, credentialPermissionsOutput+"2")

	credentialNameOutput := terraform.Output(t, terraformOptions, "credential_name")
	assert.Equal(t, credential_name, credentialNameOutput)

	credentialTopicOutput := terraform.Output(t, terraformOptions, "credential_topic")
	assert.Equal(t, kafka_credential_topic_name, credentialTopicOutput)

}

func kafkaCredentialOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_credential",
		Vars: map[string]interface{}{
			"email":   email,
			"api_key": apikey,

			"cluster_name": credential_cluster_name,
			"region":       credential_cluster_region,
			"multizone":    credential_cluster_multizone,

			"topic_name":       kafka_credential_topic_name,
			"partitions":       kafka_credential_topic_partitions,
			"retention_time":   kafka_credential_topic_retention_time,
			"retention_size":   kafka_credential_topic_retention_size,
			"max_message_size": kafka_credential_topic_max_message_size,
			"cleanup_policy":   kafka_credential_topic_cleanup_policy,

			"credential_permissions": credential_permissions,
			"credential_name":        credential_name,
		},
	})

	return terraformOptions
}
