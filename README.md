# Upstash Terraform Provider

This is the repository of official [Upstash Terraform provider](https://registry.terraform.io/providers/upstash/upstash/latest).

## Installation

```hcl

terraform {
  required_providers {
    upstash = {
      source = "upstash/upstash"
      version = "x.x.x"
    }
  }
}

provider "upstash" {
  email = var.email
  api_key  = var.api_key
}
```

`email` is your registered email in Upstash.

`api_key` can be generated from Upstash Console. For more information please check our [docs](https://docs.upstash.com/howto/developerapi).

## Create Database Using Terraform

Here example code snippet that creates database:

```hcl
resource "upstash_database" "mydb" {
  database_name = "mydb3"
  region = "eu-west-1"
  tls = "true"
  multi_zone = "false"
}
```
You can check full spec and [doc from here](https://registry.terraform.io/providers/upstash/upstash/latest/docs).

## Support, Bugs Reports, Feature Requests

If you need support then you can ask your questions Upstash Team in [upstash.com](https://upstash.com) chat widget.

There is also discord channel available for community. [Please check here](https://docs.upstash.com/help/support) for more information.


## Requirements

* Terraform v0.13 and above
* Go 1.16 (to build the provider)

## Development

If you want to locally build/test the provider then follow these steps:

* Build the provider using: `go build .` command, it will create executable in same directory
* create `terraform.rc` file that contains following configuration.
* export `TF_CLI_CONFIG_FILE` env variable that locates `terraform.rc` file.
* Now your `terraform` commands will use local Upstash provider. 
```hcl
provider_installation {

  dev_overrides {
    "upstash" = "[PATH THAT CONTAINS CUSTOM PROVIDER]"
  }
  direct {}
}
```