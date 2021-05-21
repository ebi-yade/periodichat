resource "aws_dynamodb_table" "this" {
  name           = "job-${local.pj}"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "UUID"
  range_key      = "Divisor"

  attribute {
    name = "UUID"
    type = "S"
  }

  attribute {
    name = "Divisor"
    type = "N"
  }
}
