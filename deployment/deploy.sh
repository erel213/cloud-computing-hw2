# build the new image to the ecr
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
REGION=$(aws configure get region)
REPOSITORY_NAME="whatsapp-like-repo"


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

# Force the new image to be deployed