---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "upstash_kafka_topic Resource - terraform-provider-upstash"
subcategory: ""
description: |-
  
---

# upstash_kafka_topic (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **cleanup_policy** (String) Cleanup policy will be used in the topic(compact or delete)
- **cluster_id** (String) ID of the cluster the topic will be deployed in
- **max_message_size** (Number) Max message size in the topic
- **partitions** (Number) The number of partitions the topic will have
- **retention_size** (Number) Retention size of the messages in the topic
- **retention_time** (Number) Retention time of messsages in the topic
- **topic_name** (String) Name of the topic

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **encoded_username** (String) Base64 encoded username to be used in rest communication
- **multizone** (Boolean) Whether kafka topic has multizone attribute
- **password** (String, Sensitive) Password to be used in authenticating to the cluster
- **region** (String) Region of the kafka topic
- **rest_endpoint** (String) REST Endpoint of the kafka topic
- **state** (String) State of the kafka topic (active or deleted)
- **tcp_endpoint** (String) TCP Endpoint of the kafka topic
- **topic_id** (String) Unique Cluster ID for created topic
- **username** (String) Username to be used in authenticating to the cluster

