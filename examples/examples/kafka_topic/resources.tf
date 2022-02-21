resource "upstash_kafka_cluster" "exampleKafkaCluster" {
    cluster_name = "Terraform_Upstash_Cluster"
    region = "eu-west-1"
    multizone = false
}


resource "upstash_kafka_topic" "exampleKafkaTopic" {
    topic_name = "TerraformTopic"
    partitions = 1
    retention_time = 625135
    retention_size = 725124
    max_message_size = 829213
    cleanup_policy = "delete"
    cluster_id = resource.upstash_kafka_cluster.exampleKafkaCluster.cluster_id
}

