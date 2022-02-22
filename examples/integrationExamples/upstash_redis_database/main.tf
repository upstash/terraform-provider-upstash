terraform {
  required_providers {
    upstash = {
      source = "upstash/upstash"
      version = "0.0.1"
    }
  }
}

provider "upstash" {
  email = var.email
  api_key = var.api_key
}

resource "upstash_redis_database" "exampleDB" {
  database_name = var.database_name
  region = var.region
  tls = var.tls
  multizone = var.multizone
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