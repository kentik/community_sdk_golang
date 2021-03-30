terraform {
  required_providers {
    kentik-cloudexport = {
      version = "0.1.0"
      source  = "kentik.com/automation/kentik-cloudexport"
    }
  }
}

provider "kentik-cloudexport" {
  # email, token and apiurl privided in KTAPI_AUTH_EMAIL, KTAPI_AUTH_TOKEN, KTAPI_URL env variables

  # email="john@acme.com"
  # token="token123"
  # apiurl = "http://localhost:8080"
}

# create cloudexport for AWS
resource "kentik-cloudexport_item" "terraform_aws_export" {
  name           = "test_terraform_aws_export"
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

# create cloudexport for AZURE
resource "kentik-cloudexport_item" "terraform_azure_export" {
  name           = "test_terraform_azure_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = true
  description    = "terraform azure cloud export"
  plan_id        = "11467"
  cloud_provider = "azure"
  azure {
    location                   = "centralus"
    resource_group             = "traffic-generator"
    storage_account            = "kentikstorage"
    subscription_id            = "414bd5ec-122b-41b7-9715-22f23d5b49c8"
    security_principal_enabled = true

  }
}

output "azure" {
  value = kentik-cloudexport_item.terraform_azure_export
}

# create cloudexport for IBM
resource "kentik-cloudexport_item" "terraform_ibm_export" {
  name           = "test_terraform_ibm_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = false
  description    = "terraform ibm cloud export"
  plan_id        = "11467"
  cloud_provider = "ibm"
  ibm {
    bucket = "terraform-ibm-bucket"
  }
}

output "ibm" {
  value = kentik-cloudexport_item.terraform_ibm_export
}

# create cloudexport for GCE
resource "kentik-cloudexport_item" "terraform_gce_export" {
  name           = "test_terraform_gce_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = false
  description    = "terraform gce cloud export"
  plan_id        = "11467"
  cloud_provider = "gce"
  gce {
    project      = "project gce"
    subscription = "subscription gce"
  }
}

output "gce" {
  value = kentik-cloudexport_item.terraform_gce_export
}

# create cloudexport for BGP
resource "kentik-cloudexport_item" "terraform_bgp_export" {
  name           = "test_terraform_bgp_export"
  type           = "CLOUD_EXPORT_TYPE_KENTIK_MANAGED"
  enabled        = false
  description    = "terraform bgp cloud export2"
  plan_id        = "11467"
  cloud_provider = "bgp"
  bgp {
    apply_bgp         = true
    use_bgp_device_id = "1324"
    device_bgp_type   = "router"
  }
}

output "bgp" {
  value = kentik-cloudexport_item.terraform_bgp_export
}