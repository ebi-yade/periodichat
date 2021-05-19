resource "aws_api_gateway_account" "this" {
  cloudwatch_role_arn = aws_iam_role.this.arn
}

resource "aws_iam_role" "this" {
  name               = "api_gateway_cloudwatch_global"
  assume_role_policy = data.aws_iam_policy_document.assume.json
}

data "aws_iam_policy_document" "assume" {
  statement {
    actions = [
    "sts:AssumeRole"]

    principals {
      type = "Service"
      identifiers = [
      "apigateway.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy" "this" {
  name = "default"
  role = aws_iam_role.this.id

  policy = data.aws_iam_policy_document.this.json
}

# the same as a managed policy: AmazonAPIGatewayPushToCloudWatchLogs
data "aws_iam_policy_document" "this" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:DescribeLogGroups",
      "logs:DescribeLogStreams",
      "logs:PutLogEvents",
      "logs:GetLogEvents",
      "logs:FilterLogEvents",
    ]
    resources = [
    "*"]
  }
}
