variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "private_app_subnet_cidrs" {
  description = "CIDR blocks for private application subnets"
  type        = list(string)
  default     = ["10.0.3.0/24", "10.0.4.0/24"]
}

variable "private_data_subnet_cidrs" {
  description = "CIDR blocks for private data subnets"
  type        = list(string)
  default     = ["10.0.5.0/24", "10.0.6.0/24"]
}

variable "azs" {
  description = "Availability zones"
  type        = list(string)
  default     = ["eu-west-1a", "eu-west-1b"]
}

variable "allowed_ssh_cidr" {
  description = "CIDR block allowed to SSH into the bastion host"
  type        = string
  default     = "0.0.0.0/0"
}

variable "db_allocated_storage" {
  description = "Allocated storage for the RDS instance"
  type        = number
  default     = 20
}

variable "db_engine_version" {
  description = "Engine version for the RDS instance"
  type        = string
  default     = "16.3"
}

variable "db_instance_class" {
  description = "Instance class for the RDS instance"
  type        = string
  default     = "db.t3.micro"
}

variable "db_name" {
  description = "Database name"
  type        = string
  default     = "mydatabase"
}

variable "cluster_name" {
  description = "Name of the ECS cluster"
  type        = string
  default     = "whatsappLikeCluster"
}

variable "family" {
  description = "Name of the task definition family"
  type        = string
  default     = "whatsapp-like-task"
}

variable "cpu" {
  description = "CPU units for the task"
  type        = string
  default     = "256"
}

variable "memory" {
  description = "Memory for the task"
  type        = string
  default     = "512"
}

variable "container_name" {
  description = "Name of the container"
  type        = string
  default     = "whatsapp-like-container"
}

variable "container_image" {
  description = "Container image to use"
  type        = string
}

variable "container_port" {
  description = "Port the container is listening on"
  type        = number
  default     = 8080
}

variable "service_name" {
  description = "Name of the ECS service"
  type        = string
  default     = "whatsapp-like-service"
}

variable "desired_count" {
  description = "Desired number of tasks"
  type        = number
  default     = 2
}

variable "bastion_ami" {
  description = "AMI ID for the bastion host"
  type        = string
  default     = "ami-0551ce4d67096d606"
}

variable "bastion_instance_type" {
  description = "Instance type for the bastion host"
  type        = string
  default     = "t2.micro"
}

variable "bastion_key_name" {
  description = "Key pair name for SSH access"
  type        = string
}

variable "bastion_private_key_path" {
  description = "Path to the private key file for SSH access"
  type        = string
}
