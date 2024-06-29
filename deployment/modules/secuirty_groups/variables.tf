variable "vpc_id" {
    description = "VPC ID"
    type        = string
}

variable "allowed_ssh_cidr" {
  description = "CIDR block allowed to SSH into the bastion host"
  type        = string
}
