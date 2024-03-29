---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "upstash_kafka_credential_data Data Source - terraform-provider-upstash"
subcategory: ""
description: |-
  
---

# upstash_kafka_credential_data (Data Source)



## Example Usage

```terraform
data "upstash_kafka_credential_data" "kafkaCredentialData" {
    credential_id = upstash_kafka_credential.exampleKafkaCredential.credential_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credential_id` (String) Unique ID of the kafka credential

### Read-Only

- `cluster_id` (String) ID of the kafka cluster
- `creation_time` (Number) Creation time of the credential
- `credential_name` (String) Name of the kafka credential
- `id` (String) The ID of this resource.
- `password` (String) Password to be used in authenticating to the cluster
- `permissions` (String) Permission scope given to the kafka credential
- `state` (String) State of the credential(active or deleted)
- `topic` (String) Name of the kafka topic
- `username` (String) Username to be used for the kafka credential
