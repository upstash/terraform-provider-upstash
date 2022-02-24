data "upstash_kafka_cluster_data" "clusterData" {
  cluster_id = resource.upstash_kafka_cluster.exampleCluster.cluster_id
}

