resource "aws_dynamodb_table" "mylearn_grades" {
  name         = "mylearn-grades"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserId"
  range_key    = "ClassYear"

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "ClassYear"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}