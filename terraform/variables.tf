
variable "region" {
  type        = string
  default     = "eu-central-1"
  description = "Region of ALB Logs stack"
}

variable "deployment" {
  type        = string
  default     = "test"
  description = "Deployment type"
}