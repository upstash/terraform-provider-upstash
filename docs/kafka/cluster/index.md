---
page_title: "Upstash Kafka Cluster Provider"
description: |-
  Upstash Serverless Kafka Cluster provider is used to interact with Upstash API to manage kafka clusters
---

# Upstash Kafka Cluster Provider

Upstash Serverless Kafka Cluster provider is used to interact with Upstash API to manage kafka clusters.

Use the navigation to the left read about the available resources.

## Example Usage

```hcl
terraform {
  required_providers {
    upstash = {
      source = "upstash/upstash"
      version = "X.X.X"
    }
  }
}

# Configure the Upstash Provider
provider "upstash" {
  email = var.email
  api_key  = var.api_key
}

#Create the resources

resource "upstash_kafka_cluster" "exampleCluster" {
  cluster_name = "TerraformCluster"
  region = "eu-west-1"
  multizone = false
}
```

## Authenticate the Provider

The Upstash provider requires `email` and `api_key` in order to authenticate.
Email is your email address while registering Upstash.
API KEY can be generated from the [Upstash Console](https://console.upstash.com).
Plese see our [API KEY documentation](https://docs.upstash.com/howto/developerapi).

### Required

- **api_key** (String, Sensitive)
- **email** (String)
