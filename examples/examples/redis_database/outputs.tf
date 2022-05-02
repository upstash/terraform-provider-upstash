
output "database_name" {
  value = data.upstash_redis_database_data.exampleDBData.database_name
}

output "region" {
  value = data.upstash_redis_database_data.exampleDBData.region
}

output "multizone" {
  value = resource.upstash_redis_database.exampleDB.multizone
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

output "password" {
  value = resource.upstash_redis_database.exampleDB
  sensitive = true
}

output "port" {
  value = data.upstash_redis_database_data.exampleDBData.port
}

output "rest_token" {
  value = data.upstash_redis_database_data.exampleDBData.rest_token
  sensitive = true
}

output "read_only_rest_token" {
  value = data.upstash_redis_database_data.exampleDBData.read_only_rest_token
}