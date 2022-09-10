output "rest_api_url" {
  value = aws_api_gateway_stage.mylearn.invoke_url
}

output "websocket_chat_api_url" {
  value = aws_apigatewayv2_stage.chat_api.invoke_url
}

output "websocket_lesson_api_url" {
  value = aws_apigatewayv2_stage.mylearn_lessons_api.invoke_url
}

output "s3_api_url" {
  value = aws_api_gateway_deployment.mylearn_s3.invoke_url
}

output "cloudfront_url" {
  value = aws_cloudfront_distribution.mylearn.domain_name
}