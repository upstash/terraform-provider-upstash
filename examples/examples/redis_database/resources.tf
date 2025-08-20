resource "upstash_redis_database" "exampleDB" {
  database_name = var.database_name
  region        = var.region
  tls           = var.tls
  auto_scale    = var.auto_scale
  eviction      = var.eviction

  // below ones only work if the region is given as "global"
  primary_region = var.primary_region
  read_regions   = var.read_regions
}

# resource "upstash_redis_database" "importDB" {}