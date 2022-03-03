data "upstash_redis_database_data" "exampleDBData" {
    database_id = resource.upstash_redis_database.exampleDB.database_id
}