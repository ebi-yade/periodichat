resource "aws_lambda_function" "auth" {
  function_name = "auth-${local.pj}"
  role = aws_iam_role.lambda.arn
  handler = "auth"
  filename = data.archive_file.auth.output_path
  source_code_hash = data.archive_file.auth.output_base64sha256
  #timeout = 60
  runtime = "go1.x"
}

data "archive_file" "auth" {
  type = "zip"
  source_file = "${path.module}/../_dist/auth"
  output_path = "${path.module}/archive/auth.zip"
}

resource "aws_api_gateway_resource" "auth" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  parent_id = aws_api_gateway_rest_api.this.root_resource_id
  path_part = "auth"
}

resource "aws_lambda_permission" "auth" {
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.auth.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_api_gateway_rest_api.this.execution_arn}/*/*/${aws_api_gateway_resource.auth.path_part}"
}

# The resources below increase as the number of HTTP methods you support of this resource.
resource "aws_api_gateway_method" "auth_get" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.auth.id
  http_method = "GET"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_method_response" "auth_get" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.auth.id
  http_method = aws_api_gateway_method.auth_get.http_method
  status_code = "200"
  response_models = {
    "application/json" = "Empty"
  }

  depends_on = [
    aws_api_gateway_method.auth_get]
}

resource "aws_api_gateway_integration" "auth_get" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.auth.id
  http_method = aws_api_gateway_method.auth_get.http_method
  integration_http_method = local.lambda_function_integration_method
  type = "AWS_PROXY"
  uri = aws_lambda_function.auth.invoke_arn
}
