resource "aws_s3_bucket" "mylearn_materials" {
  bucket = format("%s-%s", "mylearn-materials", var.region)

  tags = {
    AppName = "mylearn-app"
  }
}

resource "aws_s3_bucket_acl" "mylearn_materials" {
  bucket = aws_s3_bucket.mylearn_materials.id
  acl    = "private"
}

resource "aws_s3_bucket_public_access_block" "mylearn_materials" {
  bucket = aws_s3_bucket.mylearn_materials.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}