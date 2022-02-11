# terraform {
#   required_providers {
#     upstash = {
#       source = "upstash/upstash"
#       version = "0.0.1"
#     }
#   }
# }

provider "upstash" {
  api_key = "9e09722d-530f-4696-8397-9675ad33fb60"
  email = "bylmaz744@gmail.com"
}

resource "upstash_cluster" "exampleCluster" {
  cluster_name = "TerraformCluster"
  region = "eu-west-1"
  multi_zone = false
}

data "upstash_cluster_data" "clusterData" {
  cluster_id = resource.upstash_cluster.exampleCluster.cluster_id
}

