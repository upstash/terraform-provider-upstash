// DB outputs

output "database_name" {
  value = data.upstash_redis_database_data.databaseData.database_name
}
