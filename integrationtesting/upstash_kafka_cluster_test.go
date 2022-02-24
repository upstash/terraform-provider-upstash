package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUpstashKafkaCluster(t *testing.T) {
	// t.Parallel()

	envVars := GetEnvVars()

	var (
		email        = envVars.Email
		apikey       = envVars.Apikey
		cluster_name = envVars.KafkaClusterName
		region       = envVars.KafkaClusterRegion
		multizone    = envVars.KafkaClusterMultiZone
	)

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
}
