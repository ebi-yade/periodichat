resource "aws_lambda_function" "register" {
  function_name = "register-${local.pj}"
  role = aws_iam_role.lambda.arn
  handler = "register"
  filename = data.archive_file.register.output_path
  source_code_hash = data.archive_file.register.output_base64sha256
  #timeout = 60
  runtime = "go1.x"
}

data "archive_file" "register" {
  type = "zip"
  source_file = "${path.module}/../_dist/register"
  output_path = "${path.module}/archive/register.zip"
}

resource "aws_api_gateway_resource" "register" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  parent_id = aws_api_gateway_rest_api.this.root_resource_id
  path_part = "register"
}

resource "aws_lambda_permission" "register" {
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.register.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_api_gateway_rest_api.this.execution_arn}/*/*/${aws_api_gateway_resource.register.path_part}"
}

# The resources below increase as the number of HTTP methods you support of this resource.
resource "aws_api_gateway_method" "register_post" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.register.id
  http_method = "POST"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_method_response" "register_post" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.register.id
  http_method = aws_api_gateway_method.register_post.http_method
  status_code = "200"
  response_models = {
    "application/json" = "Empty"
  }
  depends_on = [
    aws_api_gateway_method.register_post]
}

resource "aws_api_gateway_integration" "register_post" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.register.id
  http_method = aws_api_gateway_method.register_post.http_method
  integration_http_method = local.lambda_function_integration_method
  type = "AWS_PROXY"
  uri = aws_lambda_function.register.invoke_arn
}
