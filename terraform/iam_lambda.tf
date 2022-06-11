# pdc-alb-logs-notifications IAM
resource "aws_iam_role" "mylearn_grades" {
  name = format("%s-%s", "mylearn-grades", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "mylearn_grades" {
  name = format("%s-%s", "mylearn-grades", var.region)
  role = aws_iam_role.mylearn_grades.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" = "Allow",
        "Action" = [
          "dynamodb:PutItem",
          "dynamodb:GetItem",
          "dynamodb:Scan",
          "dynamodb:Query",
          "dynamodb:UpdateItem"
        ],
        "Resource" = aws_dynamodb_table.mylearn_grades.arn
      },
      {
        "Effect"   = "Allow",
        "Action"   = "dynamodb:ListTables",
        "Resource" = "*"
      }
    ]
  })
}