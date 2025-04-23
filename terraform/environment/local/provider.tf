terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.13.0"
    }
  }
}
provider "google" {
  project = "local-project"
  pubsub_custom_endpoint = "http://localhost:8085/v1/"
}