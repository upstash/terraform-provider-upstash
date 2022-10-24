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
  default = "eu-west-1"

}
variable "tls"{
  default = "false"
}
variable "multizone"{
  default = "false"
}
