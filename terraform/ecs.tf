
locals {
  region = var.region
  name = "mylearn-frontend-cluster"

  user_data = <<-EOT
    #!/bin/bash
    cat <<'EOF' >> /etc/ecs/ecs.config
    ECS_CLUSTER=${local.name}
    ECS_LOGLEVEL=debug
    EOF
  EOT
}


module "ecs" {
  source = "terraform-aws-modules/ecs/aws"

  cluster_name = local.name

  cluster_configuration = {
    execute_command_configuration = {
      logging = "OVERRIDE"
      log_configuration = {
        cloud_watch_log_group_name = "/aws/ecs/aws-ec2"
      }
    }
  }

  autoscaling_capacity_providers = {
    one = {
      auto_scaling_group_arn         = aws_autoscaling_group.main.arn
      managed_termination_protection = "ENABLED"

      managed_scaling = {
        maximum_scaling_step_size = 1
        minimum_scaling_step_size = 1
        status                    = "ENABLED"
        target_capacity           = 60
      }

      default_capacity_provider_strategy = {
        weight = 60
        base   = 20
      }
    }
  }

}

resource "aws_security_group" "main" {
  ingress {
    from_port       = 22
    to_port         = 22
    protocol        = "tcp"
    cidr_blocks     = ["0.0.0.0/0"]
  }

  ingress {
    from_port       = 443
    to_port         = 443
    protocol        = "tcp"
    cidr_blocks     = ["0.0.0.0/0"]
  }

  ingress {
    from_port       = 80
    to_port         = 80
    protocol        = "tcp"
    cidr_blocks     = ["0.0.0.0/0"]
  }

  ingress {
    from_port       = 8080
    to_port         = 8080
    protocol        = "tcp"
    cidr_blocks     = ["0.0.0.0/0"]
  }

  egress {
    from_port       = 0
    to_port         = 65535
    protocol        = "tcp"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}
data "aws_iam_policy_document" "main" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "main" {
  name               = "ecs-agent"
  assume_role_policy = data.aws_iam_policy_document.main.json
}

resource "aws_iam_role_policy_attachment" "main" {
  role       = aws_iam_role.main.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_instance_profile" "ecs_agent" {
  name = "ecs-agent"
  role = aws_iam_role.main.name
}

resource "aws_launch_configuration" "main" {
  image_id             = "ami-094d4d00fd7462815"
  iam_instance_profile = aws_iam_instance_profile.ecs_agent.name
  security_groups      = [aws_security_group.main.id]
  user_data            = "#!/bin/bash\necho ECS_CLUSTER=${local.name} >> /etc/ecs/ecs.config"
  instance_type        = "t2.micro"
}

data "aws_availability_zones" "availability_zones" {
  state = "available"
}

resource "aws_autoscaling_group" "main" {
  name                      = "asg"
  availability_zones = data.aws_availability_zones.availability_zones.names
  launch_configuration      = aws_launch_configuration.main.name

  desired_capacity          = 1
  min_size                  = 1
  max_size                  = 1
  health_check_grace_period = 300
  health_check_type         = "EC2"

  protect_from_scale_in = true

  lifecycle {
    ignore_changes = [load_balancers, target_group_arns]
  }
}

data "template_file" "docker_run" {
  template = <<EOF
[{
    "essential": true,
    "memory": 512,
    "name": "worker",
    "cpu": 1,
    "image": "public.ecr.aws/v0i7a4g0/mylearn-public-ecr:latest",
    "environment": [],
    "interactive": true,
    "portMappings": [
        {
          "hostPort": 8080,
          "protocol": "tcp",
          "containerPort": 8080
        }
      ],
    "environment": [
        {
          "name": "VUE_APP_BACKEND_ECS",
          "value": "something.cloud.com"
        }
      ]
}]
EOF
}

resource "aws_ecs_task_definition" "main" {
  family                = "mylearn-frontend"
  container_definitions = data.template_file.docker_run.rendered
}

resource "aws_ecs_service" "this" {
  name            = "hello_world"
  cluster         = module.ecs.cluster_id
  task_definition = aws_ecs_task_definition.main.arn

  desired_count = 1

  deployment_maximum_percent         = 100
  deployment_minimum_healthy_percent = 0
}