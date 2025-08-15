data "upstash_redis_database_data" "databaseData" {
  database_id = resource.upstash_redis_database.redis.database_id
}
