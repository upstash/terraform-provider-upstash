# terraform {
#   required_providers {
#     upstash = {
#       source = "upstash/upstash"
#       version = "0.0.1"
#     }
#   }
# }

provider "upstash" {
  api_key = "FILL_HERE"
  email = "FILL_HERE"
}

resource "upstash_kafka_cluster" "exampleCluster" {
  cluster_name = "TerraformCluster"
  region = "eu-west-1"
  multizone = false
}

data "upstash_kafka_cluster_data" "clusterData" {
  cluster_id = resource.upstash_kafka_cluster.exampleCluster.cluster_id
}

