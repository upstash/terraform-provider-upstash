variable "email" {
  default     = "<YOUR MAIL ADDRESS>"
  description = "Upstash user email"
  type        = string
}

variable "api_key" {
  default     = "<RELATED MANAGEMENT API KEY>"
  description = "Api key for the given user"
  type        = string
}

variable "name" {
  default = "terraform_search"
  type    = string
}

variable "region" {
  default = "us-central1"
  type    = string
}

variable "type" {
  default = "payg"
  type    = string
}