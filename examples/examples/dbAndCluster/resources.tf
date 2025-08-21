resource "upstash_redis_database" "redis" {
  database_name   = "Terraform_Upstash_Database"
  region          = "global"
  primary_region  = "eu-west-1"
  tls             = true
}
