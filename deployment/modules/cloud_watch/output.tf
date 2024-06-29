output "ecs_task_definition_log_group_name" {
  value = aws_cloudwatch_log_group.ecs_task_definition_log_group.name
}