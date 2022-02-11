// !!! DONT FORGET TO ADD DB VALUES !!!

output "cluster_id" {
  value = data.upstash_cluster_data.clusterData.cluster_id
}

output "cluster_name" {
  value = data.upstash_cluster_data.clusterData.cluster_name 
}

output "cluster_region" {
  value = data.upstash_cluster_data.clusterData.region 
}

output "cluster_type" {
  value = data.upstash_cluster_data.clusterData.type
}

output "cluster_multizone" {
  value = data.upstash_cluster_data.clusterData.multizone
}

output "cluster_tcp_endpoint" {
  value = data.upstash_cluster_data.clusterData.tcp_endpoint
}

output "cluster_rest_endpoint" {
  value = data.upstash_cluster_data.clusterData.rest_endpoint
}

output "cluster_state" {
  value = data.upstash_cluster_data.clusterData.state
}

output "cluster_username" {
  value = data.upstash_cluster_data.clusterData.username
}

output "cluster_encoded_username" {
  value = data.upstash_cluster_data.clusterData.encoded_username
}

output "cluster_password" {
  value = data.upstash_cluster_data.clusterData.password
  sensitive = true
}

output "cluster_max_retention_size" {
  value = data.upstash_cluster_data.clusterData.max_retention_size
}

output "cluster_max_retention_time" {
  value = data.upstash_cluster_data.clusterData.max_retention_time
}

output "cluster_max_messages_per_second" {
  value = data.upstash_cluster_data.clusterData.max_messages_per_second
}

output "cluster_max_message_size" {
  value = data.upstash_cluster_data.clusterData.max_message_size
}

output "cluster_max_partitions" {
  value = data.upstash_cluster_data.clusterData.max_partitions
}

// DB outputs

output "database_name" {
  value = data.upstash_database_data.databaseData.database_name
}
