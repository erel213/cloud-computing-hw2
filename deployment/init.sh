#!/bin/bash

# Function to wait for the key pair to become available
wait_for_key_pair() {
  local key_name=$1
  local max_attempts=10
  local attempt=0

  while [ $attempt -lt $max_attempts ]; do
    aws ec2 describe-key-pairs --key-name $key_name >/dev/null 2>&1
    if [ $? -eq 0 ]; then
      return 0
    fi
    echo "Waiting for key pair '$key_name' to become available..."
    attempt=$((attempt + 1))
    sleep 10
  done

  return 1
}

KEY_NAME="bastion_host_key"
KEY_PATH="$(pwd)/$KEY_NAME.pem"

# Create the key pair and save the private key
aws ec2 create-key-pair --key-name $KEY_NAME --query 'KeyMaterial' --key-format pem --output text > $KEY_PATH

# Set the correct permissions for the private key file
chmod 400 $KEY_PATH

# Check if the key was created successfully
if [ $? -eq 0 ]; then
  echo "SSH key pair '$KEY_NAME' created successfully."
  echo "Private key saved to: $KEY_PATH"
else
  echo "Failed to create SSH key pair."
  exit 1
fi

# Wait for the key pair to become available
wait_for_key_pair $KEY_NAME
if [ $? -ne 0 ]; then
  echo "Key pair '$KEY_NAME' did not become available in time."
  exit 1
fi

# Log in to Amazon ECR
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
REGION=$(aws configure get region)
REPOSITORY_NAME="whatsapp-like-repo"  # Change this to your repository name

# Create the ECR repository
aws ecr create-repository --repository-name $REPOSITORY_NAME

# Check if the repository was created successfully
if [ $? -eq 0 ]; then
  echo "ECR repository '$REPOSITORY_NAME' created successfully."
else
  echo "Failed to create ECR repository."
  exit 1
fi

aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com

# Check if the login was successful
if [ $? -eq 0 ]; then
  echo "Logged in to Amazon ECR successfully."
else
  echo "Failed to log in to Amazon ECR."
  exit 1
fi

# Build the Docker image
DOCKERFILE_PATH="../Dockerfile"
IMAGE_NAME="whatsapp-like"  # Change this to your image name

docker build --platform linux/amd64 -t $IMAGE_NAME -f $DOCKERFILE_PATH ..

# Check if the image was built successfully
if [ $? -eq 0 ]; then
  echo "Docker image '$IMAGE_NAME' built successfully."
else
  echo "Failed to build Docker image."
  exit 1
fi

# Tag the Docker image
ECR_IMAGE_URI="$AWS_ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com/$REPOSITORY_NAME:$IMAGE_NAME"
docker tag $IMAGE_NAME $ECR_IMAGE_URI

# Push the Docker image to Amazon ECR
docker push $ECR_IMAGE_URI

# Check if the image was pushed successfully
if [ $? -eq 0 ]; then
  echo "Docker image '$IMAGE_NAME' pushed to ECR successfully."
else
  echo "Failed to push Docker image to ECR."
  exit 1
fi

# Initialize Terraform
terraform init

# Validate the Terraform configuration
terraform validate

# Plan the Terraform deployment
terraform plan -out=tfplan \
  -var "container_image=$ECR_IMAGE_URI" \
  -var "bastion_key_name=$KEY_NAME" \
  -var "bastion_private_key_path=$KEY_PATH" \

# Apply the Terraform deployment
terraform apply tfplan

# Check if the Terraform apply was successful
if [ $? -eq 0 ]; then
  echo "Terraform resources deployed successfully."
else
  echo "Failed to deploy Terraform resources."
  exit 1
fi

echo "Script execution completed."
