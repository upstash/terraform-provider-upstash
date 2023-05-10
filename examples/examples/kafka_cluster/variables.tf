variable "email" {
  description = "Upstash user email"
  default     = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default     = ""
}

variable "cluster_name"{
  default = "terraform_cluster"
}
variable "region"{
  default = "eu-west-1"
}
variable "multizone"{
  default = false
}
