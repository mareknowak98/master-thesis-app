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
  authorization = "NONE"
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_grades.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
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
  authorization = "NONE"
  http_method   = "POST"
  resource_id   = aws_api_gateway_resource.mylearn_grades.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
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
    aws_api_gateway_integration.mylearn_grades_post
  ]
}

output "rest_api_url" {
  value = aws_api_gateway_stage.mylearn.invoke_url
}


##########################################