resource "upstash_kafka_cluster" "exampleCluster" {
  cluster_name = "TerraformCluster"
  region       = "eu-west-1"
  multizone    = false
}