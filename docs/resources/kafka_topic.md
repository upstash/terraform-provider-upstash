---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "upstash_kafka_topic Resource - terraform-provider-upstash"
subcategory: ""
description: |-
  
---

# upstash_kafka_topic (Resource)



## Example Usage

```terraform
# Not necessary if the topic belongs to an already created cluster.
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
    
    # Here, you can use the newly created kafka_cluster resource (above) named exampleKafkaCluster.
    # And use its ID so that the topic binds to it.

    # Alternatively, provide the ID of an already created cluster.
    cluster_id = resource.upstash_kafka_cluster.exampleKafkaCluster.cluster_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cleanup_policy` (String) Cleanup policy will be used in the topic(compact or delete)
- `cluster_id` (String) ID of the cluster the topic will be deployed in
- `max_message_size` (Number) Max message size in the topic
- `partitions` (Number) The number of partitions the topic will have
- `retention_size` (Number) Retention size of the messages in the topic
- `retention_time` (Number) Retention time of messages in the topic
- `topic_name` (String) Name of the topic

### Read-Only

- `creation_time` (Number) Creation time of the topic
- `id` (String) The ID of this resource.
- `multizone` (Boolean) Whether multizone replication is enabled
- `password` (String, Sensitive) Password to be used in authenticating to the cluster
- `region` (String) Region of the kafka topic
- `rest_endpoint` (String) REST Endpoint of the kafka topic
- `state` (String) State of the kafka topic (active or deleted)
- `tcp_endpoint` (String) TCP Endpoint of the kafka topic
- `topic_id` (String) Unique Cluster ID for created topic
- `username` (String) Base64 encoded username to be used in authenticating to the cluster
