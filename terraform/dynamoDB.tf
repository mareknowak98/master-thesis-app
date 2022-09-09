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
  hash_key     = "Username"

  attribute {
    name = "Username"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_dynamodb_table" "chat_connections" {
  name         = "chat-connections"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "userId"

  attribute {
    name = "userId"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_dynamodb_table" "chat_messages" {
  name         = "mylearn-chat-messages"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserFromTo"
  range_key    = "Timestamp"

  attribute {
    name = "UserFromTo"
    type = "S"
  }

  attribute {
    name = "Timestamp"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_dynamodb_table" "mylearn_classes" {
  name         = "mylearn-classes"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserClass"

  attribute {
    name = "UserClass"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_dynamodb_table" "mylearn_lessons_connections" {
  name         = "mylearn-lessons-connections"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LessonId"
  range_key    = "ConnectionID"

  attribute {
    name = "LessonId"
    type = "S"
  }

  attribute {
    name = "ConnectionID"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_dynamodb_table" "mylearn_lessons" {
  name         = "mylearn-lessons"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LessonId"
  range_key    = "SlideId"

  attribute {
    name = "LessonId"
    type = "S"
  }

  attribute {
    name = "SlideId"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_dynamodb_table" "mylearn_tests" {
  name         = "mylearn-tests"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "TestId"
  range_key    = "CombinedKey"

  attribute {
    name = "TestId"
    type = "S"
  }

  attribute {
    name = "CombinedKey"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_dynamodb_table" "mylearn_tests_results" {
  name         = "mylearn-tests-results"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "TestId"
  range_key    = "UserId"

  attribute {
    name = "TestId"
    type = "S"
  }

  attribute {
    name = "UserId"
    type = "S"
  }

  tags = {
    AppName = "mylearn-app-change-tag"
  }
}