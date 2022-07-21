locals {
  s3_origin_id = format("%s-%s", "myS3Origin", var.region)
}

data "aws_cloudfront_cache_policy" "caching_disabled" {
  name = "Managed-CachingDisabled"
}

resource "aws_cloudfront_distribution" "s3_distribution" {

  enabled = true
  default_cache_behavior {
    allowed_methods        = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods         = ["HEAD", "GET"]
    target_origin_id       = aws_lb.mylearn.dns_name
    viewer_protocol_policy = "redirect-to-https"
    cache_policy_id        = data.aws_cloudfront_cache_policy.caching_disabled.id
  }
  origin {
    domain_name = aws_lb.mylearn.dns_name
    origin_id   = aws_lb.mylearn.dns_name

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "http-only"
      origin_ssl_protocols   = ["TLSv1"]
    }
  }
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
  viewer_certificate {
    cloudfront_default_certificate = true
  }
}

resource "aws_s3_bucket" "host_bucket" {
  bucket = "mylearn-host-bucket"

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_s3_bucket_acl" "b_acl" {
  bucket = aws_s3_bucket.host_bucket.id
  acl    = "private"
}


