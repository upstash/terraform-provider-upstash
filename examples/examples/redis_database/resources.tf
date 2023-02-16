resource "upstash_redis_database" "exampleDB" {
  database_name = var.database_name
  region = var.region
  tls = var.tls
  auto_scale = var.auto_scale
  eviction = var.eviction
}