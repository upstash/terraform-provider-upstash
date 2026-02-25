resource "upstash_redis_database" "exampleDB" {
  database_name  = "Terraform DB6"
  platform       = "aws"
  primary_region = "eu-west-1"
  tls            = true
}