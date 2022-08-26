# mylearn-grades IAM
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


# cognito-users IAM
resource "aws_iam_role" "cognito_after_register" {
  name = format("%s-%s", "cognito-after-register", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "cognito_after_register" {
  name = format("%s-%s", "cognito-after-register", var.region)
  role = aws_iam_role.cognito_after_register.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect"   = "Allow",
        "Action"   = "dynamodb:PutItem",
        "Resource" = aws_dynamodb_table.cognito_users.arn
      },
      {
        "Effect" : "Allow",
        "Action" : "cognito-idp:*",
        "Resource" : aws_cognito_user_pool.mylearn.arn
      }
    ]
  })
}

# auth-lambda IAM
resource "aws_iam_role" "auth_lambda" {
  name = format("%s-%s", "auth-lambda", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}


# chat-websocket-lambda IAM
resource "aws_iam_role" "mylearn_chat" {
  name = format("%s-%s", "mylearn-chat", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "mylearn_chat" {
  name = format("%s-%s", "mylearn-chat", var.region)
  role = aws_iam_role.mylearn_chat.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" = "Allow",
        "Action" = [
          "dynamodb:DescribeReservedCapacityOfferings",
          "dynamodb:ListGlobalTables",
          "dynamodb:ListTables",
          "dynamodb:DescribeReservedCapacity",
          "dynamodb:ListBackups",
          "dynamodb:PurchaseReservedCapacityOfferings",
          "dynamodb:DescribeLimits",
          "dynamodb:ListStreams"
        ],
        "Resource" = "*"
      },
      {
        "Effect"   = "Allow",
        "Action"   = "dynamodb:*",
        "Resource" = aws_dynamodb_table.chat_connections.arn
      },
      {
        "Effect"   = "Allow",
        "Action"   = "dynamodb:PutItem",
        "Resource" = aws_dynamodb_table.chat_messages.arn
      },
      {
        "Effect"   = "Allow",
        "Action"   = "execute-api:*",
        "Resource" = format("%s%s", aws_apigatewayv2_api.chat_api.execution_arn, "/*/*/*")
      }
    ]
  })
}

# chat-rest-lambda IAM
resource "aws_iam_role" "mylearn_rest_chat" {
  name = format("%s-%s", "mylearn-rest-chat", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "mylearn_rest_chat" {
  name = format("%s-%s", "mylearn-rest-chat", var.region)
  role = aws_iam_role.mylearn_rest_chat.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect"   = "Allow",
        "Action"   = "dynamodb:Query",
        "Resource" = aws_dynamodb_table.chat_messages.arn
      }
    ]
  })
}


# user-lambda IAM
resource "aws_iam_role" "mylearn_users" {
  name = format("%s-%s", "user-lambda", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "mylearn_users" {
  name = format("%s-%s", "user-lambda", var.region)
  role = aws_iam_role.mylearn_users.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" = "Allow",
        "Action" = [
          "dynamodb:PutItem",
          "dynamodb:DeleteItem",
          "dynamodb:Scan",
          "dynamodb:UpdateItem",
        ],
        "Resource" = aws_dynamodb_table.mylearn_classes.arn
      },
      {
        "Effect" = "Allow",
        "Action" = [
          "cognito-idp:ListUsersInGroup",
          "cognito-idp:ListUsers",
        ],
        "Resource" = "*"
      }
    ]
  })
}

# cognito-user-lambda IAM
resource "aws_iam_role" "mylearn_cognito_user" {
  name = format("%s-%s", "cognito-user-lambda", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "mylearn_cognito_user" {
  name = format("%s-%s", "cognito-user-lambda", var.region)
  role = aws_iam_role.mylearn_cognito_user.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" = "Allow",
        "Action" = [
          "cognito-idp:InitiateAuth",
          "cognito-idp:AdminConfirmSignUp",
          "cognito-idp:AdminAddUserToGroup"
        ]
        "Resource" = "*"
      }
    ]
  })
}


#lesson-rest-lambda IAM
resource "aws_iam_role" "mylearn_rest_lessons" {
  name = format("%s-%s", "lesson-rest-lambda", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "mylearn_rest_lessons" {
  name = format("%s-%s", "lesson-rest-lambda", var.region)
  role = aws_iam_role.mylearn_rest_lessons.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" = "Allow",
        "Action" = [
          "dynamodb:PutItem",
          "dynamodb:DeleteItem",
          "dynamodb:Scan",
          "dynamodb:UpdateItem",
        ],
        "Resource" = aws_dynamodb_table.mylearn_lessons.arn
      }
    ]
  })
}


# chat-lessons-lambda IAM
resource "aws_iam_role" "mylearn_lessons" {
  name = format("%s-%s", "mylearn-lessons", var.region)

  assume_role_policy  = file("files/AWSLambdaTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

resource "aws_iam_role_policy" "mylearn_lessons" {
  name = format("%s-%s", "mylearn-lessons", var.region)
  role = aws_iam_role.mylearn_lessons.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" = "Allow",
        "Action" = [
          "dynamodb:DescribeReservedCapacityOfferings",
          "dynamodb:ListGlobalTables",
          "dynamodb:ListTables",
          "dynamodb:DescribeReservedCapacity",
          "dynamodb:ListBackups",
          "dynamodb:PurchaseReservedCapacityOfferings",
          "dynamodb:DescribeLimits",
          "dynamodb:ListStreams"
        ],
        "Resource" = "*"
      },
      {
        "Effect" = "Allow",
        "Action" = [
          "dynamodb:Query",
          "dynamodb:PutItem",
          "dynamodb:DeleteItem",
          "dynamodb:Scan"
        ]
        "Resource" = [
          aws_dynamodb_table.mylearn_lessons_connections.arn,
          aws_dynamodb_table.mylearn_lessons.arn
        ]
      },
      {
        "Effect"   = "Allow",
        "Action"   = "execute-api:*",
        "Resource" = format("%s%s", aws_apigatewayv2_api.mylearn_lessons_api.execution_arn, "/*/*/*")
      }
    ]
  })
}

#s3-content-lambda IAM
resource "aws_iam_role" "s3_content_gateway" {
  name = format("%s-%s", "s3-content-gateway", var.region)

  assume_role_policy  = file("files/ApiGatewayTrustPolicy.json")
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonAPIGatewayPushToCloudWatchLogs"]
}

resource "aws_iam_role_policy" "s3_content_gateway" {
  name = format("%s-%s", "s3-content-gateway", var.region)
  role = aws_iam_role.s3_content_gateway.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" = "Allow",
        "Action" = [
          "s3:PutObject",
          "s3:GetObject"
        ],

        "Resource" = format("%s%s", aws_s3_bucket.mylearn_materials.arn, "/*")
      }
    ]
  })
}

