resource "kentik-cloudexport_item" "terraform_aws_export" {
  name           = "demo_terraform_aws_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = true
  description    = "terraform aws cloud export"
  plan_id        = "11467"
  cloud_provider = "aws"
  aws {
    bucket            = "terraform-aws-bucket"
    iam_role_arn      = "arn:aws:iam::003740049406:role/trafficTerraformIngestRole"
    region            = "us-east-2"
    delete_after_read = false
    multiple_buckets  = false
  }
}

output "aws" {
  value = kentik-cloudexport_item.terraform_aws_export
}
