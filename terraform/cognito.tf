# Cognito user pool
resource "aws_cognito_user_pool" "mylearn" {
  name = "mylearn-cognito-userpool"

  # only admin can create users - temporary
  #  admin_create_user_config {
  #    allow_admin_create_user_only = true
  #  }

  alias_attributes = ["email"]


  # days to reset default password given by admin
  password_policy {
    minimum_length                   = 8
    temporary_password_validity_days = 100
  }

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  lambda_config {
    post_confirmation = aws_lambda_function.cognito_after_register.arn
  }

}

# Cognito user groups
resource "aws_cognito_user_group" "admins" {
  name         = "admin-group"
  user_pool_id = aws_cognito_user_pool.mylearn.id
  description  = "Group of app admins"
  precedence   = 1
}

resource "aws_cognito_user_group" "teachers" {
  name         = "teacher-group"
  user_pool_id = aws_cognito_user_pool.mylearn.id
  description  = "Group of teachers"
  precedence   = 2
}

resource "aws_cognito_user_group" "parents" {
  name         = "parent-group"
  user_pool_id = aws_cognito_user_pool.mylearn.id
  description  = "Group of student's parents"
  precedence   = 3
}

resource "aws_cognito_user_group" "students" {
  name         = "student-group"
  user_pool_id = aws_cognito_user_pool.mylearn.id
  description  = "Group of students"
  precedence   = 4
}

# Cognito client - associated with webapp-userpool
resource "aws_cognito_user_pool_client" "webapp" {
  name            = "terraform-cognito-client"
  user_pool_id    = aws_cognito_user_pool.mylearn.id
  generate_secret = false

  # allow all types of auth
  explicit_auth_flows = [
    "ALLOW_USER_SRP_AUTH",
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_ADMIN_USER_PASSWORD_AUTH",
    "ALLOW_CUSTOM_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]

  # time of token life
  refresh_token_validity = 10

  access_token_validity = 24


}