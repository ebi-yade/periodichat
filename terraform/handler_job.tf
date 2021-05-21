resource "aws_lambda_function" "job" {
  function_name    = "job-${local.pj}"
  role             = aws_iam_role.lambda.arn
  handler          = "job"
  filename         = data.archive_file.job.output_path
  source_code_hash = data.archive_file.job.output_base64sha256
  #timeout = 60
  runtime = "go1.x"

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.this.name
    }
  }
}

data "archive_file" "job" {
  type        = "zip"
  source_file = "${path.module}/../_dist/job"
  output_path = "${path.module}/archive/job.zip"
}

resource "aws_api_gateway_resource" "job" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  parent_id   = aws_api_gateway_rest_api.this.root_resource_id
  path_part   = "job"
}

resource "aws_lambda_permission" "job" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.job.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.this.execution_arn}/*/*/${aws_api_gateway_resource.job.path_part}"
}

# The resources below increase as the number of HTTP methods you support of this resource.
resource "aws_api_gateway_method" "job_post" {
  rest_api_id      = aws_api_gateway_rest_api.this.id
  resource_id      = aws_api_gateway_resource.job.id
  http_method      = "POST"
  authorization    = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_method_response" "job_post" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.job.id
  http_method = aws_api_gateway_method.job_post.http_method
  status_code = "200"
  response_models = {
    "application/json" = "Empty"
  }
  depends_on = [
  aws_api_gateway_method.job_post]
}

resource "aws_api_gateway_integration" "job_post" {
  rest_api_id             = aws_api_gateway_rest_api.this.id
  resource_id             = aws_api_gateway_resource.job.id
  http_method             = aws_api_gateway_method.job_post.http_method
  integration_http_method = local.lambda_function_integration_method
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.job.invoke_arn
}
