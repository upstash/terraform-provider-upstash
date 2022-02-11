package integration

import (
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformResourceGroup(t *testing.T) {

	options := &terraform.Options{
		TerraformDir: "../test/my_examples",
	}

	defer terraform.Destroy(t, options)

	_, err := terraform.InitE(t, options)
	if err != nil {
		log.Println(err)
	}

	_, err = terraform.PlanE(t, options)
	if err != nil {
		log.Println(err)
	}

	_, err = terraform.ApplyE(t, options)
	if err != nil {
		log.Println(err)
	}

}
