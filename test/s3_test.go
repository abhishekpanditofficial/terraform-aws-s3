package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformS3Example(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/s3",
		Vars: map[string]interface{}{
			"Client":"test",
			"bucket_name": "test-primary-terraform",	
		},
	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	// Primary S3 Test
	primaryBucket := terraform.Output(t, terraformOptions, "bucket_id")
	assert.Equal(t, "test-primary-terraform", primaryBucket)
}