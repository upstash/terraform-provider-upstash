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
  default = "terraform_index"
  type    = string
}

variable "similarity_function" {
  default = "COSINE"
  type    = string
}

variable "dimension_count" {
  default = 1536
  type    = number
}

variable "region" {
  default = "us-east-1"
  type    = string
}

variable "type" {
  default = "payg"
  type    = string
}