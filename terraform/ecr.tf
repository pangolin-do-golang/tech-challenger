 module "ecr" {
   source = "terraform-aws-modules/ecr/aws"

   repository_name = "tech-challenge"

   registry_policy        = jsonencode({
     Version = "2012-10-17",
     Statement = [
       {
         Sid    = "testpolicy",
         Effect = "Allow",
         Principal = {
           "AWS" : "arn:aws:iam::${var.account_id}:root"
         },
         Action = [
           "ecr:ReplicateImage"
         ],
         Resource = [
           "arn:aws:ecr:us-east-1:${var.account_id}:repository/*"
         ]
       }, {
         Sid    = "Registry",
         Effect = "Allow",
         Principal = {
           "AWS" : "arn:aws:iam::${var.account_id}:root"
         },
         Action = [
           "ecr:CreateRepository",
           "ecr:BatchImportUpstreamImage"
         ],
         Resource = [
           "arn:aws:ecr:us-east-1:${var.account_id}:repository/tech-challenge/*"
         ]
       }
     ]
   })

   repository_lifecycle_policy = jsonencode({
     rules = [
       {
         rulePriority = 2,
         description  = "Images Docker",
         selection = {
           tagStatus     = "tagged",
           tagPrefixList = ["v"],
           countType     = "imageCountMoreThan",
           countNumber   = 30
         },
         action = {
           type = "expire"
         }
       }
     ]
   })

   tags = {
     managedBy = "terraform"
     owner = var.owner
   }

 }