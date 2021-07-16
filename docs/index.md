---
page_title: "Upstash Provider"
subcategory: ""
description: |-
  Upstash Serverless Redis Database provider is used to interact Upstash API to manage redis databases
---

# Upstash Provider

Upstash Serverless Redis database provider is used to interact Upstash API to manage redis databases.

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

resource "upstash_database" "my_database" {
  database_name = "my_redis"
  region = "eu-west-1"
  tls = "true"
  multi_zone = "false"
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
