output "rest_api_url" {
  value = aws_api_gateway_stage.mylearn.invoke_url
}

output "websocket_chat_api_url" {
  value = aws_apigatewayv2_stage.chat_api.invoke_url
}

output "cloudfront_url" {
  value = aws_cloudfront_distribution.mylearn.domain_nameq
}