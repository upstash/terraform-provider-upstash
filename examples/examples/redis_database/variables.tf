variable "email" {
  description = "Upstash user email"
  default     = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default     = ""
}

variable "database_name"{
  default = "terraform_db"
}
variable "region"{
  default = "global"
  # or for regional, pick a region. eg default="eu-west-1"

}
variable "tls"{
  default = "false"
}
variable "multizone"{
  default = "true"
}

variable "eviction"{
  default = "true"
}

variable "auto_scale"{
  default = "true"
}

# below ones only work when region="global"
variable "primary_region" {
  default = "eu-central-1"
}

variable "read_regions"{
  type = set(string)
  default = ["us-east-1"]
}
