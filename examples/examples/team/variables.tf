variable "email" {
  description = "Upstash user email"
  default     = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default     = ""
}

variable "team_name"{}
variable "copy_cc"{}
variable "team_members"{
  type = map
}

