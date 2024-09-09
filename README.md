# Folder structure

```
 ├── card-verification          # Card verification (Verifies 16 digit card number).
 ├── host-ui                    # To host UI to AWS.
    |── automation-script       # Script to confirm the UI.
    |── cfn                     # Cloudformation template to deploy UI & Cloudfront distribution.
    |── deploy-script           # To deploy SPA to S3 bucket.
 ├── src                        # Contains UI code, index.html
 ├── .env                       # Stores Bucket_name and AWS region
 └── readme.md
```

## Golang

- Go is a primary language in order to share the data model and library among lambdas
- Go 1.23 or later https://golang.org/

## Card Verification

Execute the code by running the below command

```
cd card-verification
go run ./main.go
```

### Problem

1. It must start with a , or .
1. It must contain exactly digits.
1. It must only consist of digits (-).
1. It may have digits in groups of , separated by one hyphen "-".
1. It must NOT use any other separator like ' ' , '\_', etc.
1. It must NOT have or more consecutive repeated digits.

### Solution

1. A regex expression is added to confirm the 16 digit card number.
   1. Regex confirms the allowed 16-digit pattern with the card number starting with 4,5 or 6.
   1. Regex allows combination of both 16-digit numbers eg. `4123456789123456` or `5123-4567-8912-3456`.
   1. Regex does not allow 16-digit number `61234-567-8912-3456` or `5123 - 3567 - 8912 - 3456`.
1. Consecutive digits of 4 or more are not allowed.
   1. `hasConsecutiveRepeatedDigits` removes the `-` and runs in a for loop to count the consecutive digits.
   1. Go, unlike some other languages, the backreference `\1` cannot be used in regular expressions directly for matching repeated characters.

## AWS SPA hosting

**Architecture Overview**

1. **Amazon S3**: Hosts your static website content.
2. **Amazon CloudFront**: Provides global content delivery with low latency and HTTPS support.
3. **AWS Certificate Manager (ACM)**: Manages SSL/TLS certificates for secure HTTPS connections.
4. **AWS Identity and Access Management (IAM)**: Controls access to your S3 bucket.
5. **AWS CloudFormation**: Automates the deployment of resources.
   Golang: Performs automated tests to validate the configuration.

### UI Confirmation script

Automation script is written in Golang to confirm the SPA server configuration.

```
go test -v
```

### cfn - Cloudformation template

1. File - cdn.yml

Deploy the cdn.yml.

- `CloudFrontDistribution` Sets up a CloudFront distribution with the S3 bucket as the origin. It redirects HTTP to HTTPS and uses an ACM certificate for secure connections.
- It redirects the requuests from `http` to `https` with the `ViewerProtocolPolicy`.
- `ViewerCertificate` to verify the SSL certificate.
- `Route53RecordSet` Optionally, sets up a DNS record in Route 53 to point your domain to the CloudFront distribution.
- `S3OriginAccessIdentity` Creates a CloudFront Origin Access Identity (OAI) to securely access the S3 bucket.

2. File - ssm-param.yml

- Stores the `S3BucketOaiId` created the global region.

3. File - presentation-layer.yml

- `StaticSiteBucket` Creates an S3 bucket to host your SPA. It includes website configuration to serve index.html as the default page.
- `S3BucketPolicy` Grants public read access to the bucket so CloudFront can serve the content.

## Deploy SPA

Execute the shell script deploy.sh to create and upload the SPA file to S3 bucket.

```
sh deploy.sh
```

- deploy.sh internally calls the `delete.sh` and `copy.sh`.
  - `delete.sh` delete the contents of the S3 bucket.
  - `copy.sh` syncs the content of the local src directory to S3 bucket. It also creates the S3 bucket if not already present.

## Steps to use.

1. Create an S3 Bucket: The StaticSiteBucket resource will automatically create it.
1. Deploy the Template: Use AWS CloudFormation to deploy the template. Make sure to provide the ARN of the SSL certificate from ACM.
1. Upload Content: Upload your SPA files (e.g., index.html) to the S3 bucket created by the template.
1. Update Route 53: Replace your-hosted-zone-id with your Route 53 hosted zone ID and `www.example.com` with your domain.
1. Verify Deployment: Access your SPA through the CloudFront URL or your custom domain to ensure everything is set up correctly.

## Secure the Application

1. Restrict Bucket Access
   Ensure the bucket is not publicly accessible.
   Access is only through CloudFront via the OAI.

2. Enforce HTTPS
   CloudFront is configured to redirect HTTP requests to HTTPS.

3. Use WAF (Optional)
   For additional security, you can use AWS Web Application Firewall (WAF) with CloudFront.
