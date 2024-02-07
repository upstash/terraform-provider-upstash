package integrationtesting

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

var vectorIndexName, vectorIndexType, vectorIndexRegion, vectorIndexSimilarityFunction, vectorIndexDimensionCount string

func TestUpstashVectorIndexMAIN(t *testing.T) {
	//t.Parallel()

	envVars := GetEnvVars()

	email = envVars.Email
	apikey = envVars.Apikey
	vectorIndexName = envVars.VectorIndexName
	vectorIndexType = envVars.VectorIndexType
	vectorIndexRegion = envVars.VectorIndexRegion
	vectorIndexDimensionCount = strconv.Itoa(envVars.VectorIndexDimensionCount)
	vectorIndexSimilarityFunction = envVars.VectorIndexSimilarityFunction

	terraformOptions := vectorIndexOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	vectorIndexAsserter(t, terraformOptions)

	UpstashVectorIndexRecreate(t)
	UpstashVectorIndexUpdate(t)
}

func UpstashVectorIndexRecreate(t *testing.T) {

	vectorIndexName = vectorIndexName + "Updated"
	vectorIndexType = "fixed"

	terraformOptions := vectorIndexOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	vectorIndexAsserter(t, terraformOptions)
}

func UpstashVectorIndexUpdate(t *testing.T) {

	vectorIndexDimensionCount = "512"
	vectorIndexSimilarityFunction = "COSINE"

	terraformOptions := vectorIndexOptions(t)
	terraform.Apply(t, terraformOptions)
	terraform.Plan(t, terraformOptions)
	vectorIndexAsserter(t, terraformOptions)
}

func vectorIndexAsserter(t *testing.T, terraformOptions *terraform.Options) {
	indexNameOutput := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, vectorIndexName, indexNameOutput, "The output name should be the same as the input name")

	indexTypeOutput := terraform.Output(t, terraformOptions, "type")
	assert.Equal(t, vectorIndexType, indexTypeOutput, "The output type should be the same as the input type")

	indexRegionOutput := terraform.Output(t, terraformOptions, "region")
	assert.Equal(t, vectorIndexRegion, indexRegionOutput, "The output region should be the same as the input region")

	indexDimensionCountOutput := terraform.Output(t, terraformOptions, "dimension_count")
	assert.Equal(t, vectorIndexDimensionCount, indexDimensionCountOutput, "The output dimension count should be the same as the input dimension count")

	indexSimilarityFunctionOutput := terraform.Output(t, terraformOptions, "similarity_function")
	assert.Equal(t, vectorIndexSimilarityFunction, indexSimilarityFunctionOutput, "The output similarity function should be the same as the input similarity function")
}

func vectorIndexOptions(t *testing.T) *terraform.Options {
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/examples/vector_index",
		Vars: map[string]interface{}{
			"email":               email,
			"api_key":             apikey,
			"name":                vectorIndexName,
			"type":                vectorIndexType,
			"region":              vectorIndexRegion,
			"dimension_count":     vectorIndexDimensionCount,
			"similarity_function": vectorIndexSimilarityFunction,
		},
	}
	return terraformOptions
}
