
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

output "creation_time" {
  value = data.upstash_redis_database_data.exampleDBData.creation_time
}

output "database_type" {
  value = data.upstash_redis_database_data.exampleDBData.database_type
}

output "state" {
  value = data.upstash_redis_database_data.exampleDBData.state
}

output "user_email" {
  value = data.upstash_redis_database_data.exampleDBData.user_email
}

output "db_max_clients" {
  value = data.upstash_redis_database_data.exampleDBData.db_max_clients
}

output "db_max_request_size" {
  value = data.upstash_redis_database_data.exampleDBData.db_max_request_size
}

output "db_disk_threshold" {
  value = data.upstash_redis_database_data.exampleDBData.db_disk_threshold
}

output "db_max_entry_size" {
  value = data.upstash_redis_database_data.exampleDBData.db_max_entry_size
}

output "db_memory_threshold" {
  value = data.upstash_redis_database_data.exampleDBData.db_memory_threshold
}

output "db_daily_bandwidth_limit" {
  value = data.upstash_redis_database_data.exampleDBData.db_daily_bandwidth_limit
}

output "db_max_commands_per_second" {
  value = data.upstash_redis_database_data.exampleDBData.db_max_commands_per_second
}

output "tls" {
  value = data.upstash_redis_database_data.exampleDBData.tls
}

output "consistent" {
  value = data.upstash_redis_database_data.exampleDBData.consistent
}