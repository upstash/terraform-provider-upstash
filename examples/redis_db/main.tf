terraform {
  required_providers {
    upstash = {
      source = "upstash/upstash"
      version = "0.0.1"
    }
  }
}

provider "upstash" {
  api_key = "FILL HERE"
  email = "FILL HERE"
}

resource "upstash_database" "mydb" {
  database_name = "mydb3"
  region = "eu-west-1"
  tls = "true"
  multizone = "false"
}

output "endpoint" {
  value = resource.upstash_database.mydb.endpoint
}

output "password" {
  value = resource.upstash_database.mydb.password
  sensitive = true
}