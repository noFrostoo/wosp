variable "api_key" {
  type        = string
}

variable "region" {
  description = "The region to create resources in."
  type        = string
  default     = "waw"
}

variable "plans" {
  type = map
  default = {
    "15USD"  = "vc2-2c-2gb"
    "20USD" = "vc2-2c-4gb"
    "10USD" = "vc2-1c-2gb"
  }
}

variable "plan" {
  type = string
  default = "15USD"
}

variable "nodes_quantity" {
  type = map
  default = {
    "default"  = 1
    "max" = 3
    "min" = 1
  }
}

variable "k8s_version" {
  type        = string
  default     = "v1.29.1+1"
}

variable "k8s_label" {
  type        = string
  default     = "vke-wosp-prod"
}

variable "k8s_nodes_label" {
  type        = string
  default     = "vke-wosp-nodes-prod"
}
