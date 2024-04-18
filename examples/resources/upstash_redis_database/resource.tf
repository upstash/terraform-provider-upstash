resource "upstash_redis_database" "exampleDB" {
  database_name = "Terraform DB6"
  region        = "eu-west-1"
  tls           = "true"
  multizone     = "true"
}