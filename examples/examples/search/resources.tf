resource "upstash_search" "exampleSearchResource" {
  name                = var.name
  region              = var.region
  type                = var.type
}