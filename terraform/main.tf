terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = ">= 3.41.0, < 4"
    }
  }
}

provider "aws" {}

data "aws_caller_identity" "self" {}

output "api_base_url" {
  value = aws_api_gateway_deployment.this.invoke_url
}

module "api_gateway_account" {
  source = "./modules/api_gateway_account"
  count = var.api_gateway_accounts_already_exists ? 0 : 1
}
