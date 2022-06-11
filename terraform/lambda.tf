resource "aws_lambda_function" "mylearn_grades" {
  # Find .zip file with name in 'PdcLogsNotifications*.zip' format
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
      GRADES_TABLE = "mylearn-grades"
    }
  }
}