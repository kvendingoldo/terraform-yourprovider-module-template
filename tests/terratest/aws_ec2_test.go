package terratests

//
// NOTE: This is a sample Terraform code. Delete the context of it after the GitHub project fork.
//

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAwsEC2Instance(t *testing.T) {
	// Define Terraform options to point to the Terraform folder
	terraformOptions := &terraform.Options{
		TerraformDir: "../aws_ec2_example",
		Vars: map[string]interface{}{
			"ami_id": "ami-0c55b159cbfafe1f0",
		},
	}

	// At the end of the test, destroy the resources
	defer terraform.Destroy(t, terraformOptions)

	// Init and apply the Terraform configuration
	terraform.InitAndApply(t, terraformOptions)

	// Fetch the output variable
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// Run a simple assertion to ensure the instance ID is not empty
	assert.NotEmpty(t, instanceID, "Instance ID should not be empty")
}
