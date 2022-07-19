# terraform file use only for deploying docker image to ECR

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.9.0"
    }
  }
}

provider "aws" {
  profile = "default"
  alias  = "us_east_1"
  region = "us-east-1"
}

resource "aws_ecrpublic_repository" "webapp-public-ecr" {
  # only region for ECR
  provider        = aws.us_east_1
  repository_name = "mylearn-public-ecr"

  catalog_data {
    operating_systems = ["Linux"]
  }

  # build docker image from local and push to ECR
  provisioner "local-exec" {
    command = "aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/${self.registry_id}"
  }
  provisioner "local-exec" {
    command = "docker build -t ${self.repository_name} ../"
  }
  provisioner "local-exec" {
    command = "docker tag ${self.repository_name}:latest public.ecr.aws/${self.registry_id}/${self.repository_name}:latest"
  }
  provisioner "local-exec" {
    command = "docker push public.ecr.aws/${self.registry_id}/${self.repository_name}:latest"
  }
}