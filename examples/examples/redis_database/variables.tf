variable "email" {
  description = "Upstash user email"
  default     = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default     = ""
}

variable "database_name" {
  default = "terraform_db"
}

variable "platform" {
  type    = string
  default = "aws"
}

variable "primary_region" {
  type = string
  default = "us-east-1"
}
variable "multizone" {
  default = "true"
}

variable "eviction" {
  default = "true"
}

variable "auto_scale" {
  default = "true"
}


variable "read_regions" {
  type    = set(string)
  default = ["eu-central-1"]
}
