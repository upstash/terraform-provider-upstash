variable "email" {
  default = "fahreddin@upstash.com"
  description = "Upstash user email"
  type = string
}

variable "api_key" {
  default     = "fb3b1b26-c2c0-43d9-a0fa-1a528b01b633"
  description = "Api key for the given user"
  type = string
}

variable "name" {
  default = "terraform_index_rename"
  type = string
}

variable "similarity_function" {
  default = "COSINE"
  type = string
}

variable "dimension_count" {
  default = 1536
  type = number
}

variable "region" {
  default = "us-east-1"
  type = string
}

variable "type" {
  default = "payg"
  type = string
}