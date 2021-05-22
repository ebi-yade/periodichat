resource "aws_cloudwatch_event_rule" "this" {
  name                = local.pj
  description         = "every minute"
  schedule_expression = "cron(* * * * ? *)"
}

resource "aws_cloudwatch_event_target" "this" {
  rule      = aws_cloudwatch_event_rule.this.name
  target_id = "output_report"
  arn       = aws_lambda_function.eventbridge.arn
}

resource "aws_lambda_permission" "this" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.eventbridge.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.this.arn
}
