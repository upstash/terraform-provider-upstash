resource "upstash_vector_index" "exampleVectorResource" {
  name                = var.name
  similarity_function = var.similarity_function
  dimension_count     = var.dimension_count
  region              = var.region
  type                = var.type
}