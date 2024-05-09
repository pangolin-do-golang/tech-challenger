# Create Infrastructure

## Requirements To Deploy Infrastructure

- Create a Bucket S3 to Save Backend Terraform State 
  - Add Secrets in GitHub Secrets Environment
    - AWS_ACCESS_KEY_ID `YOUR_ACCESS_KEY`
    - AWS_SECRET_ACCESS_KEY `YOUR_SECRET_ACCESS_KEY`
    - BUCKET_TF_STATE `YOUR_BUCKET_NAME`