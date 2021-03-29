terraform {
  required_providers {
    kentik-cloudexport = {
      version = "0.1.0"
      source  = "kentik.com/automation/kentik-cloudexport"
    }
  }
}

provider "kentik-cloudexport" {
  apiurl = "http://localhost:8080"

  # email and token not used when connecting to localhost_apiserver
  # email="john@acme.com"
  # token="token123"
}

# fetch all cloud exports
data "kentik-cloudexport_list" "exports" {}

output "export_list" {
  value = data.kentik-cloudexport_list.exports
}

# fetch single cloud export - it must exist or error will be reported
# data "kentik-cloudexport_item" "export" {
#   id = "3"
# }

# output "export_3" {
#   value = data.kentik-cloudexport_item.export
# }