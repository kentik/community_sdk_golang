# fetch single cloud export - it must exist or error will be reported
data "kentik-cloudexport_item" "export" {
  id = "3"
}

output "export_3" {
  value = data.kentik-cloudexport_item.export
}