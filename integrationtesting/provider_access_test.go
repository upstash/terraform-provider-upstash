package integrationtesting

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAllUpstashResources(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../test/dbAndCluster",
	}

	defer terraform.Destroy(t, terraformOptions)

	// terraform.InitAndApply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	terraform.Apply(t, terraformOptions)

	clusterName := terraform.Output(t, terraformOptions, "cluster_name")
	clusterRegion := terraform.Output(t, terraformOptions, "cluster_region")
	clusterMultizone := terraform.Output(t, terraformOptions, "cluster_multizone")
	clusterState := terraform.Output(t, terraformOptions, "cluster_state")

	assert.Equal(t, "Terraform_Upstash_Cluster", clusterName)
	assert.Equal(t, "eu-west-1", clusterRegion)
	assert.Equal(t, "false", clusterMultizone)
	assert.Equal(t, "active", clusterState)

	databaseName := terraform.Output(t, terraformOptions, "database_name")

	assert.Equal(t, "Terraform_Upstash_Database", databaseName)

}
