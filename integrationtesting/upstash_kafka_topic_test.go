package integrationtesting

import (
	"strconv"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var kafka_topic_name, kafka_topic_cleanup_policy string
var kafka_topic_partitions, kafka_topic_retention_time, kafka_topic_retention_size, kafka_topic_max_message_size int

var cluster_name, cluster_region string
var cluster_multizone bool

func TestUpstashKafkaTopicMAIN(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey

	cluster_name = envVars.KafkaClusterName
	cluster_region = envVars.KafkaClusterRegion
	cluster_multizone = envVars.KafkaClusterMultiZone

	kafka_topic_name = envVars.KafkaTopicName
	kafka_topic_partitions = envVars.KafkaTopicPartitions
	kafka_topic_retention_time = envVars.KafkaTopicRetentionTime
	kafka_topic_retention_size = envVars.KafkaTopicRetentionSize
	kafka_topic_max_message_size = envVars.KafkaTopicMaxMessageSize
	kafka_topic_cleanup_policy = envVars.KafkaTopicCleanupPolicy

	terraformOptions := kafkaTopicOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaTopicAsserter(t, terraformOptions)

	UpstashKafkaTopicRecreate(t)
	UpstashKafkaTopicUpdate(t)

}

func UpstashKafkaTopicRecreate(t *testing.T) {
	kafka_topic_name = kafka_topic_name + "Updated"
	kafka_topic_partitions = kafka_topic_partitions + 1
	kafka_topic_retention_time = kafka_topic_retention_time * 15
	kafka_topic_retention_size = kafka_topic_retention_size * 15
	// kafka_topic_max_message_size = kafka_topic_max_message_size
	// kafka_topic_cleanup_policy = kafka_topic_cleanup_policy

	terraformOptions := kafkaTopicOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaTopicAsserter(t, terraformOptions)

}

func UpstashKafkaTopicUpdate(t *testing.T) {

	kafka_topic_retention_time = kafka_topic_retention_time / 20
	kafka_topic_retention_size = kafka_topic_retention_size / 20
	kafka_topic_max_message_size = kafka_topic_max_message_size / 20

	terraformOptions := kafkaTopicOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	kafkaTopicAsserter(t, terraformOptions)

}

func kafkaTopicAsserter(t *testing.T, terraformOptions *terraform.Options) {
	// Cluster assertions
	clusterNameOutput := terraform.Output(t, terraformOptions, "cluster_name")
	assert.Equal(t, cluster_name, clusterNameOutput)

	regionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, cluster_region, regionOutput)

	multizoneOutput := terraform.Output(t, terraformOptions, "multizone") == "true"
	assert.Equal(t, cluster_multizone, multizoneOutput)

	// Topic assertions
	topicNameOutput := terraform.Output(t, terraformOptions, "topic_name")
	assert.Equal(t, kafka_topic_name, topicNameOutput)

	partitionsOutput := terraform.Output(t, terraformOptions, "partitions")
	assert.Equal(t, strconv.Itoa(kafka_topic_partitions), partitionsOutput)

	// retention_timeOutput := terraform.Output(t, terraformOptions, "retention_time")
	// assert.Equal(t, strconv.Itoa(kafka_topic_retention_time), retention_timeOutput)

	// retention_sizeOutput := terraform.Output(t, terraformOptions, "retention_size")
	// assert.Equal(t, strconv.Itoa(kafka_topic_retention_size), retention_sizeOutput)

	max_message_sizeOutput := terraform.Output(t, terraformOptions, "max_message_size")
	assert.Equal(t, strconv.Itoa(kafka_topic_max_message_size), max_message_sizeOutput)

	cleanup_policyOutput := terraform.Output(t, terraformOptions, "cleanup_policy")
	assert.Equal(t, kafka_topic_cleanup_policy, cleanup_policyOutput)
}

func kafkaTopicOptions(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/examples/kafka_topic",
		Vars: map[string]interface{}{
			"email":   email,
			"api_key": apikey,

			"cluster_name": cluster_name,
			"region":       cluster_region,
			"multizone":    cluster_multizone,

			"topic_name":       kafka_topic_name,
			"partitions":       kafka_topic_partitions,
			"retention_time":   kafka_topic_retention_time,
			"retention_size":   kafka_topic_retention_size,
			"max_message_size": kafka_topic_max_message_size,
			"cleanup_policy":   kafka_topic_cleanup_policy,
		},
	})

	return terraformOptions
}
