resource "aws_lambda_function" "mylearn_grades" {
  # Find .zip file with name in 'grades-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{grades-lambda}*.zip")))
  function_name = "mylearn-grades"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.mylearn_grades.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION       = var.region
      GRADES_TABLE = aws_dynamodb_table.mylearn_grades.name
    }
  }
}

resource "aws_lambda_permission" "cognito_after_register" {
  statement_id  = "AllowExecutionFromCognito"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cognito_after_register.function_name
  principal     = "cognito-idp.amazonaws.com"
  source_arn    = aws_cognito_user_pool.mylearn.arn
  depends_on    = [aws_lambda_function.cognito_after_register]
}

resource "aws_lambda_function" "cognito_after_register" {
  # Find .zip file with name in 'cognito-after-register-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{cognito-after-register-lambda}*.zip")))
  function_name = "cognito-after-register-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.cognito_after_register.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION       = var.region
      GRADES_TABLE = aws_dynamodb_table.cognito_users.name
    }
  }
}

resource "aws_lambda_function" "auth_lambda" {
  # Find .zip file with name in 'auth-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{auth-lambda}*.zip")))
  function_name = "auth-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.auth_lambda.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION       = var.region
      USER_POOL_ID = aws_cognito_user_pool.mylearn.id
    }
  }
}


resource "aws_lambda_function" "chat_lambda" {
  # Find .zip file with name in 'chat-websocket-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{chat-websocket-lambda}*.zip")))
  function_name = "chat-websocket-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.mylearn_chat.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION            = var.region
      USER_POOL_ID      = aws_cognito_user_pool.mylearn.id
      CONNECTIONS_TABLE = aws_dynamodb_table.chat_connections.name
      DEPLOYMENT_TYPE   = var.deployment
      MESSAGES_TABLE    = aws_dynamodb_table.chat_messages.name
      WEBSOCKET_API     = aws_apigatewayv2_api.chat_api.id
    }
  }
}

resource "aws_lambda_function" "chat_rest_lambda" {
  # Find .zip file with name in 'chat-websocket-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{chat-rest-lambda}*.zip")))
  function_name = "chat-rest-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.mylearn_rest_chat.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION         = var.region
      MESSAGES_TABLE = aws_dynamodb_table.chat_messages.name
    }
  }
}


resource "aws_lambda_function" "user_lambda" {
  # Find .zip file with name in 'user-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{user-lambda}*.zip")))
  function_name = "user-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.mylearn_users.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION            = var.region
      USER_TABLE        = aws_dynamodb_table.cognito_users.name
      USER_POOL_ID      = aws_cognito_user_pool.mylearn.id
      COGNITO_CLIENT_ID = aws_cognito_user_pool_client.webapp.id
      CLASS_TABLE       = aws_dynamodb_table.mylearn_classes.name
    }
  }
}


resource "aws_lambda_function" "cognito_user_lambda" {
  # Find .zip file with name in 'user-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{cognito-user-lambda}*.zip")))
  function_name = "cognito-user-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.mylearn_cognito_user.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION            = var.region
      USER_POOL_ID      = aws_cognito_user_pool.mylearn.id
      COGNITO_CLIENT_ID = aws_cognito_user_pool_client.webapp.id
    }
  }
}

resource "aws_lambda_function" "lessons_rest_lambda" {
  # Find .zip file with name in 'user-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{lesson-rest-lambda}*.zip")))
  function_name = "lesson-rest-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.mylearn_rest_lessons.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION        = var.region
      LESSONS_TABLE = aws_dynamodb_table.mylearn_lessons.name
    }
  }
}


resource "aws_lambda_function" "lessons_websocket_lambda" {
  # Find .zip file with name in 'chat-websocket-lambda*.zip' format
  filename      = format("%s/%s", "../lambdas/lambda_build", one(fileset("../lambdas/lambda_build", "{lesson-websocket-lambda}*.zip")))
  function_name = "lesson-websocket-lambda"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.mylearn_lessons.arn
  timeout       = 15
  memory_size   = 128

  tags = {
    AppName = "mylearn-app"
  }

  environment {
    variables = {
      REGION            = var.region
      USER_POOL_ID      = aws_cognito_user_pool.mylearn.id
      CONNECTIONS_TABLE = aws_dynamodb_table.mylearn_lessons_connections.name
      DEPLOYMENT_TYPE   = var.deployment
      LESSONS_TABLE     = aws_dynamodb_table.mylearn_lessons.name
      WEBSOCKET_API     = aws_apigatewayv2_api.mylearn_lessons_api.id
    }
  }
}