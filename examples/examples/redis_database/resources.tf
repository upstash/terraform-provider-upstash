resource "upstash_redis_database" "exampleDB" {
  database_name = "tf testa"
  region        = "global"
  tls           = true
  auto_scale    = true
  eviction      = true
  prod_pack     = true
  budget = 30

  primary_region = "eu-west-1"
}

resource "upstash_redis_database" "exampleDB" {
  database_name = var.database_name
  region        = "global"
  tls           = true
  auto_scale    = var.auto_scale
  eviction      = var.eviction
  prod_pack     = var.prod_pack
  budget        = var.budget


  primary_region = var.primary_region
  read_regions   = var.read_regions
}

# resource "upstash_redis_database" "importDB" {}