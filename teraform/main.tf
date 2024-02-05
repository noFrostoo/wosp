terraform {
  required_providers {
    vultr = {
      source = "vultr/vultr"
      version = "2.19.0"
    }
  }
}

provider "vultr" {
  api_key = var.api_key
  rate_limit = 100
  retry_limit = 3
}

resource "vultr_kubernetes" "k8" {
    region  = var.region
    label   = var.k8s_label
    version = var.k8s_version
} 


resource "vultr_kubernetes_node_pools" "np" {
    cluster_id    = vultr_kubernetes.k8.id
    node_quantity = var.nodes_quantity["default"]
    plan          = var.plans[var.plan]
    label         = var.k8s_nodes_label
    auto_scaler   = true
    min_nodes     = var.nodes_quantity["min"]
    max_nodes     = var.nodes_quantity["max"]
}