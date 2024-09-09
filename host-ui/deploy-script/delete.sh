#!/bin/bash

aws s3 rm "s3://$BUCKET_NAME" --recursive --profile $REGION
