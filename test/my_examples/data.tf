data "upstash_database_data" "databaseData" {
    database_id = resource.upstash_database.mydb.database_id
}

data "upstash_cluster_data" "clusterData" {
  cluster_id = resource.upstash_cluster.exampleCluster.cluster_id
}