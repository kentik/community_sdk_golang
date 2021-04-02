# Note: configuration for locally-built provider
terraform {
  required_providers {
    kentik-cloudexport = {
      version = "0.1.0"
      source  = "kentik.com/automation/kentik-cloudexport"
    }
  }
}

provider "kentik-cloudexport" {
  # email, token and apiurl are read from KTAPI_AUTH_EMAIL, KTAPI_AUTH_TOKEN, KTAPI_URL env variables
}