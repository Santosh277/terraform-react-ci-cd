package test

import (
  "testing"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/stretchr/testify/assert"
)

func TestTerraformBasic(t *testing.T) {
  t.Parallel()

  terraformOptions := &terraform.Options{
    TerraformDir: "../", // path to your main.tf
  }

  // Init and Apply Terraform
  defer terraform.Destroy(t, terraformOptions)
  terraform.InitAndApply(t, terraformOptions)

  // Example: Check output variable
  bucketName := terraform.Output(t, terraformOptions, "artifact_bucket_name")
  assert.Contains(t, bucketName, "artifact")
}
