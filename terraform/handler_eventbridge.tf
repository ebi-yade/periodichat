locals {
  ev_function_name = "eventbridge-${local.pj}-${substr(filemd5("../script/update-lambda-env.sh"), 0, 7)}"
}
resource "aws_lambda_function" "eventbridge" {
  function_name    = local.ev_function_name
  role             = aws_iam_role.lambda.arn
  handler          = "eventbridge"
  filename         = data.archive_file.eventbridge.output_path
  source_code_hash = data.archive_file.eventbridge.output_base64sha256
  timeout          = 60
  runtime          = "go1.x"
  environment {
    variables = {
      TABLE_NAME      = aws_dynamodb_table.this.name
      ZOOM_AUTH_TOKEN = "dummy"
    }
  }
  lifecycle {
    ignore_changes = [environment]
  }

  provisioner "local-exec" {
    command = "bash -c ../script/update-lambda-env.sh"

    environment = {
      FN_NAME = local.ev_function_name
    }
  }
}

data "archive_file" "eventbridge" {
  type        = "zip"
  source_file = "${path.module}/../_dist/eventbridge"
  output_path = "${path.module}/archive/eventbridge.zip"
}
