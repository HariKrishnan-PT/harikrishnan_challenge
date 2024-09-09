#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# Use the loaded environment variables
echo "Deploying to bucket: $BUCKET_NAME in region: $REGION"

# Ensure the bucket exists
if ! aws s3 ls "s3://${BUCKET_NAME}" > /dev/null 2>&1; then
  echo "Bucket does not exist. Creating bucket..."
  aws s3 mb "s3://${BUCKET_NAME}" --region "${REGION}"
fi

aws s3 sync ../src/ "s3://$BUCKET_NAME" --profile $REGION
