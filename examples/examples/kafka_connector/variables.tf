variable "email" {
  description = "Upstash user email"
  default     = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default     = ""
}

variable "cluster_name" {
  default = "clusterForConnector"
}
variable "region" {
  default = "eu-west-1"
}
variable "multizone" {
  default = false
}

variable "connector_name" {
  default = "test_connector2"
}
