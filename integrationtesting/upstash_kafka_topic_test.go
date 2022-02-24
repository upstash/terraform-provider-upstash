package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUpstashKafkaTopic(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	var (
		email  = envVars.Email
		apikey = envVars.Apikey

		cluster_name = envVars.KafkaClusterName
		region       = envVars.KafkaClusterRegion
		multizone    = envVars.KafkaClusterMultiZone

		topic_name       = envVars.KafkaTopicName
		partitions       = envVars.KafkaTopicPartitions
		retention_time   = envVars.KafkaTopicRetentionTime
		retention_size   = envVars.KafkaTopicRetentionSize
		max_message_size = envVars.KafkaTopicMaxMessageSize
		cleanup_policy   = envVars.KafkaTopicCleanupPolicy
	)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_topic",
		Vars: map[string]interface{}{
			"email":   email,
			"api_key": apikey,

			"cluster_name": cluster_name,
			"region":       region,
			"multizone":    multizone,

			"topic_name":       topic_name,
			"partitions":       partitions,
			"retention_time":   retention_time,
			"retention_size":   retention_size,
			"max_message_size": max_message_size,
			"cleanup_policy":   cleanup_policy,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	// Since using built provider, no need to install from the version
	// terraform.Init(t, terraformOptions)

	terraform.Apply(t, terraformOptions)

	terraform.Plan(t, terraformOptions)
}
