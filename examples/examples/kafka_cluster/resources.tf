
resource "upstash_kafka_cluster" "exampleCluster" {
  cluster_name = var.cluster_name
  region = var.region
  multizone = var.multizone
}
