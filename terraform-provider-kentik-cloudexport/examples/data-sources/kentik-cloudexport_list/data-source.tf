# fetch all cloud exports
data "kentik-cloudexport_list" "exports" {
}

output "export_list" {
  value = data.kentik-cloudexport_list.exports
}