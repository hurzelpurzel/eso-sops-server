provider "aws" {
  region = "eu-central-1"
}

# S3 Bucket
resource "aws_s3_bucket" "example_bucket" {
  bucket = "mein-example-bucket-12345"
  acl    = "private"

  tags = {
    Name        = "MeinBucket"
    Environment = "Dev"
  }
}

# IAM User
resource "aws_iam_user" "bucket_user" {
  name = "bucket-user"
}

# IAM Policy für Zugriff auf den Bucket
resource "aws_iam_policy" "bucket_policy" {
  name        = "bucket-access-policy"
  description = "Policy für Zugriff auf den S3 Bucket"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["s3:ListBucket"]
        Effect   = "Allow"
        Resource = aws_s3_bucket.example_bucket.arn
      },
      {
        Action   = ["s3:GetObject", "s3:PutObject", "s3:DeleteObject"]
        Effect   = "Allow"
        Resource = "${aws_s3_bucket.example_bucket.arn}/*"
      }
    ]
  })
}

# Policy an den User anhängen
resource "aws_iam_user_policy_attachment" "attach_bucket_policy" {
  user       = aws_iam_user.bucket_user.name
  policy_arn = aws_iam_policy.bucket_policy.arn
}

# Access Keys für den User
resource "aws_iam_access_key" "bucket_user_key" {
  user = aws_iam_user.bucket_user.name
}

# Outputs
output "bucket_name" {
  value = aws_s3_bucket.example_bucket.bucket
}

output "iam_user_name" {
  value = aws_iam_user.bucket_user.name
}

output "access_key_id" {
  value     = aws_iam_access_key.bucket_user_key.id
  sensitive = true
}

output "secret_access_key" {
  value     = aws_iam_access_key.bucket_user_key.secret
  sensitive = true
}
