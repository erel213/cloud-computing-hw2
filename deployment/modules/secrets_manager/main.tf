# Create kms
resource "aws_kms_key" "secrets_manager" {
  description             = "KMS key for secrets manager"
  deletion_window_in_days = 10
  is_enabled = true
  enable_key_rotation = true

  tags = {
    Name = "kms-secrets-manager"
  }
}

# Create random password
resource "random_password" "db_credentials_password" {
  length           = 16
  special          = true
  override_special = "_%"
}

resource "aws_secretsmanager_secret" "db-password" {
  kms_key_id = aws_kms_key.secrets_manager.key_id
  name = "db-password"
  description = "RDS Admin Password"
  recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "db-password" {
  secret_id = aws_secretsmanager_secret.db-password.id
  secret_string = random_password.db_credentials_password.result
}

resource "aws_secretsmanager_secret" "ecs_rds_connection_string" {
  name = "rds-connection-string"
  description = "Connection string for RDS for ECS tasks"
  recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "ecs_rds_connection_string" {
  secret_id = aws_secretsmanager_secret.ecs_rds_connection_string.id
  secret_string = "postgresql://root:${aws_secretsmanager_secret_version.db-password.secret_string}@${var.rds_endpoint}/whatsapp-like"
}