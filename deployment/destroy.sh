# Remove bastion key
aws ec2 delete-key-pair \
--key-name bastion_host_key

rm -f bastion_host_key.pem

# Remove ECR repository
aws ecr delete-repository \
--repository-name whatsapp-like-repo \
--force