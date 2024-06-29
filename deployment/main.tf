provider "aws" {
  region = "eu-west-1"
}


module "vpc" {
  source                    = "./modules/vpc"
  vpc_cidr                  = var.vpc_cidr
  public_subnet_cidrs       = var.public_subnet_cidrs
  private_app_subnet_cidrs  = var.private_app_subnet_cidrs
  private_data_subnet_cidrs = var.private_data_subnet_cidrs
  azs                       = var.azs
}

module "security_groups" {
  source           = "./modules/secuirty_groups"
  vpc_id           = module.vpc.vpc_id
  allowed_ssh_cidr = var.allowed_ssh_cidr
}

module cloud_watch {
  source = "./modules/cloud_watch"
}

module "load_balancer" {
  source            = "./modules/load_balancer"
  vpc_id            = module.vpc.vpc_id
  public_subnet_ids = module.vpc.public_subnet_ids
  security_group_id = module.security_groups.lb_sg_id
}

module "secrets_manager" {
  source = "./modules/secrets_manager"
  rds_endpoint = module.rds.rds_endpoint
}

module "rds" {
  source                  = "./modules/rds"
  vpc_id                  = module.vpc.vpc_id
  private_data_subnet_ids = module.vpc.private_data_subnet_ids
  allocated_storage       = var.db_allocated_storage
  engine_version          = var.db_engine_version
  instance_class          = var.db_instance_class
  security_group_id       = module.security_groups.db_sg_id
  db_credentials_secret_arn  = module.secrets_manager.db_credentials_secret_arn
}

module "ecs" {
  source           = "./modules/ecs"
  cluster_name     = var.cluster_name
  family           = var.family
  cpu              = var.cpu
  memory           = var.memory
  container_name   = var.container_name
  container_image  = var.container_image
  container_port   = var.container_port
  service_name     = var.service_name
  desired_count    = var.desired_count
  subnet_ids       = module.vpc.private_app_subnet_ids
  app_sg_id        = module.security_groups.app_sg_id
  target_group_arn = module.load_balancer.lb_target_group_arn
  db_credentials_secret_arn  = module.secrets_manager.db_credentials_secret_arn
  depends_on = [ module.secrets_manager, module.rds ]
  aws_cloudwatch_log_group_name = module.cloud_watch.ecs_task_definition_log_group_name
  connection_string_arn = module.secrets_manager.connection_string_arn
}

module "bastion" {
  source           = "./modules/bastion"
  ami              = var.bastion_ami
  instance_type    = var.bastion_instance_type
  public_subnet_id = element(module.vpc.public_subnet_ids, 0)
  vpc_id           = module.vpc.vpc_id
  key_name         = var.bastion_key_name
  private_key_path = var.bastion_private_key_path
  allowed_ssh_cidr = var.allowed_ssh_cidr
  app_sg_id        = module.security_groups.app_sg_id
  bastion_sg_id    = module.security_groups.bastion_sg_id
}
