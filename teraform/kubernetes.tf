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