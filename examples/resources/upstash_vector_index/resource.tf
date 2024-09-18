resource "upstash_vector_index" "vectorResource" {
  name                = "vectorResource"
  similarity_function = "COSINE"
  dimension_count     = 1536
  region              = "us-east-1"
  type                = "fixed"
}