data "upstash_redis_database_data" "databaseData" {
  database_id = resource.upstash_redis_database.redis.database_id
}

data "upstash_kafka_cluster_data" "clusterData" {
  cluster_id = resource.upstash_kafka_cluster.exampleCluster.cluster_id
}