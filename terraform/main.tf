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

output "account" {
  value = data.aws_caller_identity.self.account_id
}
