variable "email" {
  description = "Upstash user email"
  default = ""
}
variable "api_key" {
  description = "Api key for the given user"
  default = ""
}

variable "cluster_name"{
  default = "clusterForCredential"
}
variable "region"{
  default = "eu-west-1" 
}
variable "multizone"{
  default = false
}

variable "topic_name"{
  default = "topicForCredential"
}
variable "partitions"{
  default = 2
}
variable "retention_time"{
  default = 100000
}
variable "retention_size"{
  default = 100000
}
variable "max_message_size"{
  default = 100000
}
variable "cleanup_policy"{
  default = "delete"
}

variable "credential_name"{}
variable "credential_permissions"{}
