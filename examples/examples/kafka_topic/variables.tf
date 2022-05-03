variable "email" {
  description = "Upstash user email"
  default = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default = ""
}

variable "cluster_name"{}
variable "region"{}
variable "multizone"{}


variable "topic_name"{}
variable "partitions"{}
variable "retention_time"{}
variable "retention_size"{}
variable "max_message_size"{}
variable "cleanup_policy"{}


