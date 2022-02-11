// May need to tweak here.

resource "upstash_database" "mydb" {
  database_name = "Terraform_Upstash_Database"
  region = "eu-west-1"
  tls = true
  multizone = false
  consistent = true
}

resource "upstash_cluster" "exampleCluster" {
  cluster_name = "Terraform_Upstash_Cluster"
  region = "eu-west-1"
  multizone = false
}