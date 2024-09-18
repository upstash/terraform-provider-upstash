resource "upstash_kafka_cluster" "exampleKafkaCluster" {
  cluster_name = var.cluster_name
  region       = var.region
  multizone    = var.multizone
}


resource "upstash_kafka_topic" "exampleKafkaTopic" {
  topic_name       = var.topic_name
  partitions       = var.partitions
  retention_time   = var.retention_time
  retention_size   = var.retention_size
  max_message_size = var.max_message_size
  cleanup_policy   = var.cleanup_policy
  cluster_id       = resource.upstash_kafka_cluster.exampleKafkaCluster.cluster_id

  # Or, if you defined locals, you can use those as such 
  # retention_size = local.retention_size_bytes
  # max_message_size = local.max_message_size_bytes
}


# resource "upstash_kafka_topic" "importKafkaTopic" {}

