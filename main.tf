provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_namespace" "notes" {
  metadata {
    name = "notes"
  }
}

module "k8s_resources" {
  source = "./k8s"

  namespace = kubernetes_namespace.notes.metadata.0.name
}
