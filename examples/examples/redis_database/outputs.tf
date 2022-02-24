
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