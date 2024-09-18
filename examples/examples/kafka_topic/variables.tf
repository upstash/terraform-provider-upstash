variable "email" {
  description = "Upstash user email"
  default     = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default     = ""
}

variable "cluster_name" {}
variable "region" {}
variable "multizone" {}


variable "topic_name" {}
variable "partitions" {}
variable "retention_time" {}
variable "retention_size" {}
variable "max_message_size" {}
variable "cleanup_policy" {}


# OR, you can use locals to convert higher numbers to bytes
# locals {
#   max_message_size_mb = 1
#   max_message_size_bytes = local.max_message_size_mb * 1024 * 1024
# }

# locals {
#   retention_size_gb = 100
#   retention_size_bytes = local.retention_size_gb * 1024 * 1024 * 1024
# }