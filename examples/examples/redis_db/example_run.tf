# terraform {
  required_providers {
    upstash = {
      source = "upstash/upstash"
      version = "X.X.X"
    }
  }
}

provider "upstash" {
  api_key = "FILL_HERE"
  email = "FILL_HERE"
}

resource "upstash_redis_database" "exampleDB" {
  database_name = "Terraform DB6"
  region = "eu-west-1"
  tls = "true"
  multizone = "true"
}

data "upstash_redis_database_data" "exampleDBData" {
    database_id = resource.upstash_redis_database.exampleDB.database_id
}

output "DBName" {
  value = data.upstash_redis_database_data.exampleDBData.database_name
}

output "DBEndpoint" {
  value = data.upstash_redis_database_data.exampleDBData.endpoint
}

output "endpoint" {
  value = resource.upstash_redis_database.exampleDB.endpoint
}

output "db_name" {
  value = resource.upstash_redis_database.exampleDB.database_name
}

output "multizone" {
  value = resource.upstash_redis_database.exampleDB.multizone
}

output "password" {
  value = resource.upstash_redis_database.exampleDB
  sensitive = true
}