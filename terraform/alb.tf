## data block referring to default VPC in giver region
#data "aws_vpc" "default" {
#  default = true
#}
#
## data block referring ids of subnets of default VPC
#data "aws_subnet_ids" "default_subnets_ids" {
#  vpc_id = data.aws_vpc.default.id
#}
#
## data block with list of default subnets
#data "aws_subnet" "default_subnet_list" {
#  count = length(data.aws_subnet_ids.default_subnets_ids.ids)
#  id    = tolist(data.aws_subnet_ids.default_subnets_ids.ids)[count.index]
#}
#
#resource "aws_lb_target_group" "mylearn" {
#  name        = "tf-example-lb-alb-tg"
#  target_type = "instance"
#  port        = 80
#  protocol    = "TCP"
#  vpc_id      = data.aws_vpc.default.id
#
#  health_check {
#    interval            = 30
#    port                = 80
#    healthy_threshold   = 5
#    unhealthy_threshold = 5
#    protocol            = "TCP"
#  }
#}
#
#resource "aws_lb" "mylearn" {
#  name               = "mylearn-alb"
#  internal           = false
#  load_balancer_type = "network"
#  subnets            = data.aws_subnet.default_subnet_list.*.id
#
#  enable_deletion_protection = false
#
#  tags = {
#    AppName = "mylearn-app"
#  }
#}
#
#resource "aws_lb_listener" "mylearn" {
#  load_balancer_arn = aws_lb.mylearn.arn
#  port              = "80"
#  protocol          = "TCP"
#
#  default_action {
#    type             = "forward"
#    target_group_arn = aws_lb_target_group.mylearn.arn
#  }
#}