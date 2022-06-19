# Budget plan quotas
resource "aws_budgets_budget" "mylearn" {
  name              = "mylearn-monthly-budget"
  budget_type       = "COST"
  limit_amount      = "2.0"
  limit_unit        = "USD"
  time_unit         = "MONTHLY"
  time_period_start = "2020-02-20_00:01"

  notification {
    comparison_operator        = "GREATER_THAN"
    threshold                  = 50
    threshold_type             = "PERCENTAGE"
    notification_type          = "FORECASTED"
    subscriber_email_addresses = ["mark-now@wp.pl", "mark.now997@gmail.com", "marek.nowak31254@gmail.com"]
  }
}
