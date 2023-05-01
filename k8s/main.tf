variable "namespace" {}

resource "kubernetes_deployment" "notes-api" {
  metadata {
    name      = "notes-api"
    namespace = var.namespace
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "notes-api"
      }
    }

    template {
      metadata {
        labels = {
          app = "notes-api"
        }
      }

      spec {
        container {
          name  = "notes-api"
          image = "algharib/golang-notes-api:latest"  # Updated image

          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "notes-api" {
  metadata {
    name      = "notes-api"
    namespace = var.namespace
  }

  spec {
    selector = {
      app = "notes-api"
    }

    type = "LoadBalancer"

    port {
      protocol    = "TCP"
      port        = 80
      target_port = 8080
    }
  }
}
