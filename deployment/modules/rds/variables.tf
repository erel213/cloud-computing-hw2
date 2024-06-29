variable "vpc_id" {
  description = "VPC ID"
  type        = string
}

variable "private_data_subnet_ids" {
  description = "Private data subnet IDs"
  type        = list(string)
}

variable "allocated_storage" {
  description = "Allocated storage for the RDS instance"
  type        = number
  default     = 20
}

variable "engine_version" {
  description = "Engine version for the RDS instance"
  type        = string
  default     = "13.4"
}

variable "instance_class" {
  description = "Instance class for the RDS instance"
  type        = string
  default     = "db.t3.micro"
}

variable "security_group_id" {
  type = string
}

variable "db_credentials_secret_arn" {
  description = "The ARN of the Secrets Manager secret for DB credentials"
  type        = string
}