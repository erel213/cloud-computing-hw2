resource "aws_instance" "bastion" {
  ami           = var.ami
  instance_type = var.instance_type
  subnet_id     = var.public_subnet_id
  key_name      = var.key_name

  security_groups = [var.bastion_sg_id]

  tags = {
    Name = "bastion-host"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
    ]


    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file(var.private_key_path)
      host        = self.public_ip
    }
  }
}