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

module "cors3" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_users.id
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

module "cors5" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_rest_chat.id
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

resource "aws_api_gateway_method_response" "mylearn_cognito_users_login_post_200" {
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
  resource_id = aws_api_gateway_resource.mylearn_cognito_users_login.id
  http_method = aws_api_gateway_method.mylearn_cognito_users_login_post.http_method
  status_code = "200"
  response_models = {
    "application/json" = "Empty"
  }
}

resource "aws_api_gateway_integration" "mylearn_cognito_users_login_post" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_cognito_users_login.id
  http_method             = aws_api_gateway_method.mylearn_cognito_users_login_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.cognito_user_lambda.invoke_arn
}


module "cors" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_cognito_users_login.id
}

#####
# cognito users management service endpoints
resource "aws_api_gateway_resource" "mylearn_cognito_users_register" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "register"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}

resource "aws_api_gateway_method" "mylearn_cognito_users_register_post" {
  http_method   = "POST"
  resource_id   = aws_api_gateway_resource.mylearn_cognito_users_register.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "mylearn_cognito_users_register_post_200" {
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
  resource_id = aws_api_gateway_resource.mylearn_cognito_users_register.id
  http_method = aws_api_gateway_method.mylearn_cognito_users_register_post.http_method
  status_code = "200"
  response_models = {
    "application/json" = "Empty"
  }
}

resource "aws_api_gateway_integration" "mylearn_cognito_users_register_post" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_cognito_users_register.id
  http_method             = aws_api_gateway_method.mylearn_cognito_users_register_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.cognito_user_lambda.invoke_arn
}


module "cors2" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_cognito_users_register.id
}


#####

resource "aws_lambda_permission" "mylearn_cognito_users" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cognito_user_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mylearn.execution_arn}/*/*"
}


#########################################
#classes
########################################
# chat service endpoints
resource "aws_api_gateway_resource" "mylearn_classes" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "classes"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}
#
resource "aws_api_gateway_method" "mylearn_classes_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_classes.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_classes_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_classes.id
  http_method             = aws_api_gateway_method.mylearn_classes_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}
#
resource "aws_api_gateway_method" "mylearn_classes_post" {
  http_method   = "POST"
  resource_id   = aws_api_gateway_resource.mylearn_classes.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_classes_post" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_classes.id
  http_method             = aws_api_gateway_method.mylearn_classes_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}
#
#
resource "aws_api_gateway_method" "mylearn_classes_delete" {
  http_method   = "DELETE"
  resource_id   = aws_api_gateway_resource.mylearn_classes.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_classes_delete" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_classes.id
  http_method             = aws_api_gateway_method.mylearn_classes_delete.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}
#

module "cors6" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_classes.id
}


##########################################
#### classUsers endpoints
resource "aws_api_gateway_resource" "mylearn_class_users" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "classUsers"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}
#
resource "aws_api_gateway_method" "mylearn_class_users_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_class_users.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_class_users_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_class_users.id
  http_method             = aws_api_gateway_method.mylearn_class_users_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}
#
resource "aws_api_gateway_method" "mylearn_class_users_post" {
  http_method   = "POST"
  resource_id   = aws_api_gateway_resource.mylearn_class_users.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_class_users_post" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_class_users.id
  http_method             = aws_api_gateway_method.mylearn_class_users_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}
#
#
resource "aws_api_gateway_method" "mylearn_class_users_delete" {
  http_method   = "DELETE"
  resource_id   = aws_api_gateway_resource.mylearn_class_users.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_class_users_delete" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_class_users.id
  http_method             = aws_api_gateway_method.mylearn_class_users_delete.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}
#

module "cors7" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_class_users.id
}


##########################################
#### me endpoints
resource "aws_api_gateway_resource" "mylearn_me" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "me"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}
#
resource "aws_api_gateway_method" "mylearn_me_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_me.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_me_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_me.id
  http_method             = aws_api_gateway_method.mylearn_me_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.user_lambda.invoke_arn
}

#
module "cors8" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_me.id
}



##########################################
# chat service endpoints
resource "aws_api_gateway_resource" "mylearn_rest_lessons" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "slides"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}

#
resource "aws_api_gateway_method" "mylearn_rest_lessons_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_rest_lessons.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_rest_lessons_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_rest_lessons.id
  http_method             = aws_api_gateway_method.mylearn_rest_lessons_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lessons_rest_lambda.invoke_arn
}
#

#
resource "aws_api_gateway_method" "mylearn_rest_lessons_post" {
  http_method   = "POST"
  resource_id   = aws_api_gateway_resource.mylearn_rest_lessons.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_rest_lessons_post" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_rest_lessons.id
  http_method             = aws_api_gateway_method.mylearn_rest_lessons_post.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lessons_rest_lambda.invoke_arn
}
#

#
resource "aws_api_gateway_method" "mylearn_rest_lessons_delete" {
  http_method   = "DELETE"
  resource_id   = aws_api_gateway_resource.mylearn_rest_lessons.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_rest_lessons_delete" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_rest_lessons.id
  http_method             = aws_api_gateway_method.mylearn_rest_lessons_delete.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lessons_rest_lambda.invoke_arn
}
#

resource "aws_lambda_permission" "mylearn_rest_lessons" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lessons_rest_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.mylearn.execution_arn}/*/*"
}

module "cors9" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_rest_lessons.id
}



##########################################
#### s3 bucket
resource "aws_api_gateway_resource" "mylearn_s3_files" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "files"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}
#
resource "aws_api_gateway_method" "mylearn_s3_files_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_s3_files.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_s3_files_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_s3_files.id
  http_method             = aws_api_gateway_method.mylearn_s3_files_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.s3_management.invoke_arn
}

#
module "cors10" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_s3_files.id
}


resource "aws_api_gateway_resource" "mylearn_s3_folders" {
  parent_id   = aws_api_gateway_rest_api.mylearn.root_resource_id
  path_part   = "folders"
  rest_api_id = aws_api_gateway_rest_api.mylearn.id
}
#
resource "aws_api_gateway_method" "mylearn_s3_folders_get" {
  http_method   = "GET"
  resource_id   = aws_api_gateway_resource.mylearn_s3_folders.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn.id
  authorization = "CUSTOM"
  authorizer_id = aws_api_gateway_authorizer.mylearn.id
}

resource "aws_api_gateway_integration" "mylearn_s3_folders_get" {
  rest_api_id             = aws_api_gateway_rest_api.mylearn.id
  resource_id             = aws_api_gateway_resource.mylearn_s3_folders.id
  http_method             = aws_api_gateway_method.mylearn_s3_folders_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.s3_management.invoke_arn
}

#
module "cors11" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn.id
  api_resource_id = aws_api_gateway_resource.mylearn_s3_folders.id
}



##########################################
##########################################
##########################################
##########################################



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
    aws_api_gateway_integration.mylearn_users_get,
    aws_api_gateway_integration.mylearn_cognito_users_login_post,
    aws_api_gateway_integration.mylearn_rest_chat_get,
    aws_api_gateway_integration.mylearn_classes_post,
    aws_api_gateway_integration.mylearn_classes_get,
    aws_api_gateway_integration.mylearn_classes_delete,
    aws_api_gateway_integration.mylearn_class_users_post,
    aws_api_gateway_integration.mylearn_class_users_get,
    aws_api_gateway_integration.mylearn_class_users_delete,
    aws_api_gateway_integration.mylearn_rest_lessons_post,
    aws_api_gateway_integration.mylearn_rest_lessons_get,
    aws_api_gateway_integration.mylearn_rest_lessons_delete,
    aws_api_gateway_integration.mylearn_s3_files_get,
    aws_api_gateway_integration.mylearn_s3_folders_get,
  ]
}



##########################################