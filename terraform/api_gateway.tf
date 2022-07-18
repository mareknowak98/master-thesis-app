# Rest API Gateway
resource "aws_api_gateway_rest_api" "mylearn" {
  name        = "mylearn-rest-api"
  description = "REST API for master thesis service"
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

##########################################
# Grades service endpoints
resource "aws_api_gateway_resource" "mylearn_grades" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "grades"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}

resource "aws_api_gateway_method" "mylearn_grades_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_grades.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_grades_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_grades.id
  http_method             = aws_api_gateway_method.mylearn_grades_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.mylearn_grades.invoke_arn
}

resource "aws_api_gateway_method" "mylearn_grades_post" {
  http_method   = "POST"
  resource_id   = aws_api_gateway_resource.mylearn_grades.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_grades_post" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_grades.id
  http_method             = aws_api_gateway_method.mylearn_grades_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.mylearn_grades.invoke_arn
}

resource "aws_lambda_permission" "mylearn_grades" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.mylearn_grades.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mylearn.execution_arn}/*/*"
}

##########################################
# users service endpoints
resource "aws_api_gateway_resource" "mylearn_users" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "users"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}

resource "aws_api_gateway_method" "mylearn_users_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_users.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_users_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_users.id
  http_method             = aws_api_gateway_method.mylearn_users_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}

resource "aws_lambda_permission" "mylearn_users" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.user_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mylearn.execution_arn}/*/*"
}

##########################################
# chat service endpoints
resource "aws_api_gateway_resource" "mylearn_rest_chat" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "messages"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}

resource "aws_api_gateway_method" "mylearn_rest_chat_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_rest_chat.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_rest_chat_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_rest_chat.id
  http_method             = aws_api_gateway_method.mylearn_rest_chat_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.chat_rest_lambda.invoke_arn
}

resource "aws_lambda_permission" "mylearn_rest_chat" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.chat_rest_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mylearn.execution_arn}/*/*"
}

##########################################
# cognito users management service endpoints
resource "aws_api_gateway_resource" "mylearn_cognito_users_login" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "login"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}

resource "aws_api_gateway_method" "mylearn_cognito_users_login_post" {
  http_method   = "POST"
  resource_id   = aws_api_gateway_resource.mylearn_cognito_users_login.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "mylearn_cognito_users_login_post" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_cognito_users_login.id
  http_method             = aws_api_gateway_method.mylearn_cognito_users_login_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.cognito_user_lambda.invoke_arn
}

resource "aws_lambda_permission" "mylearn_cognito_users" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cognito_user_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mylearn.execution_arn}/*/*"
}


##########################################
## API deployment
resource "aws_api_gateway_stage" "mylearn" {
  deployment_id = aws_api_gateway_deployment.mylearn.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  stage_name    = var.deployment
}

## deploy each run - might be added conditions
resource "aws_api_gateway_deployment" "mylearn" {
  rest_api_id = aws_api_gateway_rest_api.mylearn.id

  variables = {
    deployed_at = timestamp()
  }

  lifecycle {
    create_before_destroy = true
  }

  depends_on = [
    aws_api_gateway_integration.mylearn_grades_get,
    aws_api_gateway_integration.mylearn_grades_post,
    aws_api_gateway_integration.mylearn_users_get
  ]
}

output "rest_api_url" {
  value = aws_api_gateway_stage.mylearn.invoke_url
}


##########################################