resource "upstash_kafka_cluster" "exampleKafkaCluster" {
    cluster_name = var.cluster_name
    region = var.region
    multizone = var.multizone
}


resource "upstash_kafka_topic" "exampleKafkaTopic" {
    topic_name = var.topic_name
    partitions = var.partitions
    retention_time = var.retention_time
    retention_size = var.retention_size
    max_message_size = var.max_message_size
    cleanup_policy = var.cleanup_policy
    cluster_id = resource.upstash_kafka_cluster.exampleKafkaCluster.cluster_id
}

resource "upstash_kafka_credential" "exampleKafkaCredential" {
    cluster_id = upstash_kafka_cluster.exampleKafkaCluster.cluster_id
    credential_name=var.credential_name
    topic = upstash_kafka_topic.exampleKafkaTopic.topic_name
    permissions = var.credential_permissions
}

