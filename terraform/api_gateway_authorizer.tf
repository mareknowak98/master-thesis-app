resource "aws_api_gateway_authorizer" "mylearn" {
  name                   = "mylearn-custom-lambda-authorizer"
  rest_api_id            = aws_api_gateway_rest_api.mylearn.id
  authorizer_uri         = aws_lambda_function.auth_lambda.invoke_arn
  authorizer_credentials = aws_iam_role.mylearn_invocation_role.arn
}

resource "aws_iam_role" "mylearn_invocation_role" {
  name = "api_gateway_auth_invocation"
  path = "/"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "invocation_policy" {
  name = "default"
  role = aws_iam_role.mylearn_invocation_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "lambda:InvokeFunction",
      "Effect": "Allow",
      "Resource": "${aws_lambda_function.auth_lambda.arn}"
    }
  ]
}
EOF
}