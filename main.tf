terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.13.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "traefik" {
  name         = "traefik:latest"
  keep_locally = true
}


resource "docker_image" "golang" {
  name         = "golang:latest"
  keep_locally = true
}

resource "docker_container" "traefik" {
  name  = "traefik"
  image = docker_image.traefik.latest
  ports {
    internal = 80
    external = 8080
  }
}

resource "docker_container" "wizdwarfs" {
  name  = "wizdwarfs"
  image = var.golang_dock
  ports {
    internal = 5000
    external = 5000
  }
}
