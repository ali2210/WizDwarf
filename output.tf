output "golang_container" {
  description = "wizdwarfs container image"
  value       = docker_container.wizdwarfs.image
}

output "golang_logs" {
  value       = docker_container.wizdwarfs.logs
  description = "wizdwarfs log files"
}

output "golang_port" {
  description = " wizdwarfs port to connect"
  value       = docker_container.wizdwarfs.ports
}