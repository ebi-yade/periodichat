resource "aws_api_gateway_rest_api" "this" {
  name = local.pj
}

resource "aws_api_gateway_deployment" "this" {
  rest_api_id       = aws_api_gateway_rest_api.this.id
  stage_name        = local.pj
  stage_description = "timestamp = ${timestamp()}"

  depends_on = [
    aws_api_gateway_integration.register_post,
    aws_api_gateway_integration.auth_get,
  ]

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_method_settings" "this" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  stage_name  = aws_api_gateway_deployment.this.stage_name
  method_path = "*/*"

  settings {
    data_trace_enabled = true
    logging_level      = "ERROR"
  }

  depends_on = [
  module.api_gateway_account]
}
