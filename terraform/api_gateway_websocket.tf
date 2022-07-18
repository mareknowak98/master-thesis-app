#websocket api gateway for chat

resource "aws_apigatewayv2_api" "chat_api" {
  name                       = "mylearn-chat-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
  description                = "Websocket API for chat service"
}

# routes
resource "aws_apigatewayv2_route" "chat_connect" {
  api_id    = aws_apigatewayv2_api.chat_api.id
  route_key = "$connect"
  target    = "integrations/${aws_apigatewayv2_integration.chat_connect.id}"
}

resource "aws_apigatewayv2_integration" "chat_connect" {
  api_id                    = aws_apigatewayv2_api.chat_api.id
  integration_type          = "AWS_PROXY"
#  connection_type           = "INTERNET"
#  content_handling_strategy = "CONVERT_TO_TEXT"
  integration_method        = "POST"
  integration_uri           = aws_lambda_function.chat_lambda.invoke_arn
#  passthrough_behavior      = "WHEN_NO_MATCH"

}

resource "aws_apigatewayv2_route" "chat_disconnect" {
  api_id    = aws_apigatewayv2_api.chat_api.id
  route_key = "$disconnect"
  target    = "integrations/${aws_apigatewayv2_integration.chat_disconnect.id}"
}

resource "aws_apigatewayv2_integration" "chat_disconnect" {
  api_id                    = aws_apigatewayv2_api.chat_api.id
  integration_type          = "AWS_PROXY"
#  connection_type           = "INTERNET"
#  content_handling_strategy = "CONVERT_TO_TEXT"
  #  description               = "Lambda example"
  integration_method        = "POST"
  integration_uri           = aws_lambda_function.chat_lambda.invoke_arn
#  passthrough_behavior      = "WHEN_NO_MATCH"

}

resource "aws_apigatewayv2_route" "chat_default" {
  api_id    = aws_apigatewayv2_api.chat_api.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.chat_default.id}"
}

#lambda integrations
resource "aws_apigatewayv2_integration" "chat_default" {
  api_id                    = aws_apigatewayv2_api.chat_api.id
  integration_type          = "AWS_PROXY"
#  connection_type           = "INTERNET"
#  content_handling_strategy = "CONVERT_TO_TEXT"
  integration_method        = "POST"
  integration_uri           = aws_lambda_function.chat_lambda.invoke_arn
#  passthrough_behavior      = "WHEN_NO_MATCH"

}

#deployment
resource "aws_apigatewayv2_deployment" "chat" {
  api_id      = aws_apigatewayv2_api.chat_api.id
  description = "Example deployment"

  lifecycle {
    create_before_destroy = true
  }

  depends_on = [
    aws_apigatewayv2_route.chat_connect,
    aws_apigatewayv2_route.chat_disconnect,
    aws_apigatewayv2_route.chat_default,
  ]

}

#resource "aws_apigatewayv2_deployment" "chat_disconnect" {
#  api_id      = aws_apigatewayv2_route.chat_disconnect.api_id
#  description = "Example deployment"
#
#  lifecycle {
#    create_before_destroy = true
#  }
#}
#
#resource "aws_apigatewayv2_deployment" "chat_default" {
#  api_id      = aws_apigatewayv2_route.chat_default.api_id
#  description = "Example deployment"
#
#  lifecycle {
#    create_before_destroy = true
#  }
#}

#stages
resource "aws_apigatewayv2_stage" "chat_api" {
  api_id = aws_apigatewayv2_api.chat_api.id
  name   = var.deployment
}

# Allow the API Gateway to invoke Lambda function
resource "aws_lambda_permission" "chat_api_connect" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.chat_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.chat_api.execution_arn}/*/$connect"
}

# Allow the API Gateway to invoke Lambda function
resource "aws_lambda_permission" "chat_api_disconnect" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.chat_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.chat_api.execution_arn}/*/$disconnect"
}

# Allow the API Gateway to invoke Lambda function
resource "aws_lambda_permission" "chat_api_default" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.chat_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.chat_api.execution_arn}/*/$default"
}