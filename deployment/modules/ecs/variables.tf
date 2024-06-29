variable "cluster_name" {
  description = "Name of the ECS cluster"
  type        = string
}

variable "family" {
  description = "Name of the task definition family"
  type        = string
}

variable "cpu" {
  description = "CPU units for the task"
  type        = string
}

variable "memory" {
  description = "Memory for the task"
  type        = string
}

variable "container_name" {
  description = "Name of the container"
  type        = string
}

variable "container_image" {
  description = "Container image to use"
  type        = string
}

variable "container_port" {
  description = "Port the container is listening on"
  type        = number
}

variable "service_name" {
  description = "Name of the ECS service"
  type        = string
}

variable "desired_count" {
  description = "Desired number of tasks"
  type        = number
}

variable "subnet_ids" {
  description = "List of subnet IDs for the service"
  type        = list(string)
}

variable "app_sg_id" {
  description = "Security group ID for the application"
  type        = string
}

variable "target_group_arn" {
  description = "ARN of the load balancer target group"
  type        = string
}

variable "db_credentials_secret_arn" {
  description = "The ARN of the Secrets Manager secret for DB credentials"
  type        = string
}

variable "aws_cloudwatch_log_group_name" {
  description = "Name of the CloudWatch log group"
  type        = string
}

variable "connection_string_arn" {
  description = "The ARN of the Secrets Manager secret for the connection string"
  type        = string
}
