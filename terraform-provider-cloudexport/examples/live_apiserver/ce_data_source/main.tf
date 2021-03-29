terraform {
  required_providers {
    kentik-cloudexport = {
      version = "0.1.0"
      source  = "kentik.com/automation/kentik-cloudexport"
    }
  }
}

provider "kentik-cloudexport" {
  # by default, provider reads kentikapi credentials from env variables: KTAPI_AUTH_EMAIL and KTAPI_AUTH_TOKEN 

  # email="john@acme.com"
  # token="test123"
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