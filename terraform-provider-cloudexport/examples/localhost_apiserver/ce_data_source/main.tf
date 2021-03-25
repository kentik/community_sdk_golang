terraform {
  required_providers {
    cloudexport = {
      version = "0.1.0"
      source  = "kentik.com/automation/cloudexport"
    }
  }
}

provider "cloudexport" {
  apiurl = "http://localhost:8080"

  # email and token not used when connecting to localhost_apiserver
  # email="john@acme.com"
  # token="token123"
}

# fetch all cloud exports
data "cloudexport_list" "exports" {}

output "export_list" {
  value = data.cloudexport_list.exports
}

# fetch single cloud export - it must exist or error will be reported
# data "cloudexport_item" "export" {
#   id = "3"
# }

# output "export_3" {
#   value = data.cloudexport_item.export
# }