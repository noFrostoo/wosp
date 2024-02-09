resource "vultr_object_storage" "terraform_state" {
    cluster_id = 2
    label      = "vultr-object-storage"
}