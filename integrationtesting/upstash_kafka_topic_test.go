package integrationtesting

import (
	"strconv"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var kafkaTopicName, kafkaTopicCleanupPolicy string
var kafkaTopicPartitions, kafkaTopicRetentionTime, kafkaTopicRetentionSize, kafkaTopicMaxMessageSize int

var clusterName, clusterRegion string
var clusterMultizone bool

func TestUpstashKafkaTopicMAIN(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey

	clusterName = envVars.KafkaClusterName
	clusterRegion = envVars.KafkaClusterRegion
	clusterMultizone = envVars.KafkaClusterMultiZone

	kafkaTopicName = envVars.KafkaTopicName
	kafkaTopicPartitions = envVars.KafkaTopicPartitions
	kafkaTopicRetentionTime = envVars.KafkaTopicRetentionTime
	kafkaTopicRetentionSize = envVars.KafkaTopicRetentionSize
	kafkaTopicMaxMessageSize = envVars.KafkaTopicMaxMessageSize
	kafkaTopicCleanupPolicy = envVars.KafkaTopicCleanupPolicy

	terraformOptions := kafkaTopicOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaTopicAsserter(t, terraformOptions)

	UpstashKafkaTopicRecreate(t)
	UpstashKafkaTopicUpdate(t)

}

func UpstashKafkaTopicRecreate(t *testing.T) {
	kafkaTopicName = kafkaTopicName + "Updated"
	kafkaTopicPartitions = kafkaTopicPartitions + 1
	kafkaTopicRetentionTime = kafkaTopicRetentionTime * 15
	kafkaTopicRetentionSize = kafkaTopicRetentionSize * 15
	// kafkaTopicMaxMessageSize = kafkaTopicMaxMessageSize
	// kafkaTopicCleanupPolicy = kafkaTopicCleanupPolicy

	terraformOptions := kafkaTopicOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaTopicAsserter(t, terraformOptions)

}

func UpstashKafkaTopicUpdate(t *testing.T) {

	kafkaTopicRetentionTime = kafkaTopicRetentionTime / 20
	kafkaTopicRetentionSize = kafkaTopicRetentionSize / 20
	kafkaTopicMaxMessageSize = kafkaTopicMaxMessageSize / 20

	terraformOptions := kafkaTopicOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaTopicAsserter(t, terraformOptions)

}

func kafkaTopicAsserter(t *testing.T, terraformOptions *terraform.Options) {
	// Cluster assertions
	clusterNameOutput := terraform.Output(t, terraformOptions, "clusterName")
	assert.Equal(t, clusterName, clusterNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, clusterRegion, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, clusterMultizone, multizoneOutput)

	// Topic assertions
	topicNameOutput := terraform.Output(t, terraformOptions, "topic_name")
	assert.Equal(t, kafkaTopicName, topicNameOutput)

	partitionsOutput := terraform.Output(t, terraformOptions, "partitions")
	assert.Equal(t, strconv.Itoa(kafkaTopicPartitions), partitionsOutput)

	// retention_timeOutput := terraform.Output(t, terraformOptions, "retention_time")
	// assert.Equal(t, strconv.Itoa(kafkaTopicRetentionTime), retention_timeOutput)

	// retention_sizeOutput := terraform.Output(t, terraformOptions, "retention_size")
	// assert.Equal(t, strconv.Itoa(kafkaTopicRetentionSize), retention_sizeOutput)

	max_message_sizeOutput := terraform.Output(t, terraformOptions, "max_message_size")
	assert.Equal(t, strconv.Itoa(kafkaTopicMaxMessageSize), max_message_sizeOutput)

	cleanup_policyOutput := terraform.Output(t, terraformOptions, "cleanup_policy")
	assert.Equal(t, kafkaTopicCleanupPolicy, cleanup_policyOutput)
}

func kafkaTopicOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_topic",
		Vars: map[string]interface{}{
			"email":   email,
			"api_key": apikey,

			"clusterName": clusterName,
			"region":      clusterRegion,
			"multizone":   clusterMultizone,

			"topic_name":       kafkaTopicName,
			"partitions":       kafkaTopicPartitions,
			"retention_time":   kafkaTopicRetentionTime,
			"retention_size":   kafkaTopicRetentionSize,
			"max_message_size": kafkaTopicMaxMessageSize,
			"cleanup_policy":   kafkaTopicCleanupPolicy,
		},
	})

	return terraformOptions
}
