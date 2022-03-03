resource "upstash_redis_database" "exampleDB" {
  database_name = var.database_name
  region = var.region
  tls = var.tls
  multizone = var.multizone
}