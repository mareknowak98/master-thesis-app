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