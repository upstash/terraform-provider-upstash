resource "upstash_redis_database" "exampleDB" {
  database_name = var.database_name
  tls           = true
  auto_scale    = true
  eviction      = true
  prod_pack     = true
  budget        = 30


  primary_region = var.primary_region
  platform       = var.platform
  read_regions   = var.read_regions
}

# resource "upstash_redis_database" "importDB" {}