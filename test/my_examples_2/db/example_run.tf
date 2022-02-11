terraform {
  required_providers {
    upstash = {
      source = "upstash/upstash"
      version = "0.0.1"
    }
  }
}

provider "upstash" {
  api_key = "67e824e7-0256-4c9e-9b25-065acf61f1e7"
  email = "bylmaz744@gmail.com"
}

resource "upstash_database" "exampleDB" {
  database_name = "Terraform DB"
  region = "eu-west-1"
  tls = "true"
  multi_zone = "false"
}

output "endpoint" {
  value = resource.upstash_database.exampleDB.endpoint
}

output "db_name" {
  value = resource.upstash_database.exampleDB.database_name
}

output "multi_zone" {
  value = resource.upstash_database.exampleDB.multi_zone
}

output "password" {
  value = resource.upstash_database.exampleDB
  sensitive = true
}