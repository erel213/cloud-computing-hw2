output "db_credentials_secret_arn" {
  value = aws_secretsmanager_secret.db-password.arn
}

output "connection_string_arn" {
  value = aws_secretsmanager_secret.ecs_rds_connection_string.arn
}
