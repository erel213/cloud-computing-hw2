data "aws_secretsmanager_secret" "db_credentials" {
  arn = var.db_credentials_secret_arn
}

data "aws_secretsmanager_secret_version" "db_credentials" {
  secret_id = data.aws_secretsmanager_secret.db_credentials.id
}

resource "aws_db_subnet_group" "main" {
  name       = "main-db-subnet-group"
  subnet_ids = var.private_data_subnet_ids

  tags = {
    Name = "main-db-subnet-group"
  }
}

resource "aws_db_instance" "main" {
  allocated_storage    = var.allocated_storage
  engine               = "postgres"
  engine_version       = var.engine_version
  instance_class       = var.instance_class
  username             = "root"
  password             = data.aws_secretsmanager_secret_version.db_credentials.secret_string
  parameter_group_name = "default.postgres16"
  db_subnet_group_name = aws_db_subnet_group.main.name
  vpc_security_group_ids = [var.security_group_id]

  skip_final_snapshot = true

  tags = {
    Name = "main-rds-instance"
  }
}