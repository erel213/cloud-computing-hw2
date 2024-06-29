output "lb_arn" {
  value = aws_lb.main.arn
}

output "lb_dns_name" {
  value = aws_lb.main.dns_name
}

output "lb_target_group_arn" {
  value = aws_lb_target_group.app.arn
}
