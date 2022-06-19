resource "aws_dynamodb_table" "mylearn_grades" {
  name         = "mylearn-grades"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserId"
  range_key    = "Date"

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "Date"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}


resource "aws_dynamodb_table" "cognito_users" {
  name         = "cognito-users"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserId"

  attribute {
    name = "UserId"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}