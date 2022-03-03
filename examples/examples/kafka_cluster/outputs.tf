output "cluster_id" {
  value = data.upstash_kafka_cluster_data.clusterData.cluster_id
}

output "cluster_name" {
  value = data.upstash_kafka_cluster_data.clusterData.cluster_name 
}

output "region" {
  value = data.upstash_kafka_cluster_data.clusterData.region 
}

output "type" {
  value = data.upstash_kafka_cluster_data.clusterData.type
}

output "multizone" {
  value = data.upstash_kafka_cluster_data.clusterData.multizone
}

output "tcp_endpoint" {
  value = data.upstash_kafka_cluster_data.clusterData.tcp_endpoint
}

output "rest_endpoint" {
  value = data.upstash_kafka_cluster_data.clusterData.rest_endpoint
}

output "state" {
  value = data.upstash_kafka_cluster_data.clusterData.state
}

output "username" {
  value = data.upstash_kafka_cluster_data.clusterData.username
}

output "password" {
  value = data.upstash_kafka_cluster_data.clusterData.password
  sensitive = true
}

output "max_retention_size" {
  value = data.upstash_kafka_cluster_data.clusterData.max_retention_size
}

output "max_retention_time" {
  value = data.upstash_kafka_cluster_data.clusterData.max_retention_time
}

output "max_messages_per_second" {
  value = data.upstash_kafka_cluster_data.clusterData.max_messages_per_second
}

output "max_message_size" {
  value = data.upstash_kafka_cluster_data.clusterData.max_message_size
}

output "max_partitions" {
  value = data.upstash_kafka_cluster_data.clusterData.max_partitions
}
