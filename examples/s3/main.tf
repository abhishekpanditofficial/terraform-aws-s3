terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region     = "us-east-1"
  access_key = "AKIA2EFS3QBOGIAUHGU5"
  secret_key = "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"
}

module "terraform_aws_s3" {
  source = "../../"
  Client = var.Client
  bucket_name= var.bucket_name
}