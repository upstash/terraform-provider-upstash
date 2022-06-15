resource "upstash_redis_database" "exampleDB" {
  database_name = var.database_name
  region = var.region
  multizone = var.multizone
  tls = true
  consistent = false
}