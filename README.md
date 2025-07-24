
# React CI/CD with Terraform, AWS CodePipeline, and Terratest

This project sets up a full CI/CD pipeline to build and deploy a React application using AWS services (S3, CodePipeline, CodeBuild) and Terraform. It also includes infrastructure testing using Terratest (written in Go).

---

## 🚀 Features

- Build and deploy React apps with Vite
- Automate infrastructure with Terraform
- CI/CD pipeline using AWS CodePipeline and CodeBuild
- Host frontend on S3 with static website hosting
- Infrastructure tests with Terratest (Go)

---
## Prerequisite

Ensure your React project has buildspec and Package file in root directory

```bash
📁 src
  └── App.jsx
📁 public
📄 index.html
📄 vite.config.js
📄 package.json
📄 buildspec.yml
```

---
## 1. Create an Access Key

- Go to your aws account 
- IAM user ->create
- Give access key
- Permission(S3fullaccess, codebuild,pipeline full access){for testing :adminstration access}

---
## 2. Setup a working VM

- Create an Amazon Linux EC2 instance
- Inside that configure AWS CLI
```bash
# Install AWS CLI
sudo yum install -y awscli
# Configure AWS credentials
aws configure
```
- Install Terraform
```bash
# Add HashiCorp repo
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo

# Install Terraform
sudo yum install -y terraform

# Verify
terraform -v
```


## 3. Configure `terraform.tfvars`

Create a `terraform.tfvars` file in the project root and define the following variables:

```hcl
region           = "ap-south-1"#any region you like
artifact_bucket  = "my-react-artifact-2025"#or what ever you like
github_owner     = "........"#put your user name
github_repo      = "....."#your node app repo name
github_branch    = "main" #any other if you have
github_token     = "ghp_xxxxxxxxxxxxxxxxxxxxxxx"

```

⚠️ **Warning**: Never commit `terraform.tfvars` to GitHub — it contains sensitive information.

## 3.1. Similarly Configure `main.tf` , `variables.tf` , `output.tf` .
- you can add `buildspec.yml` (not recommended) only if you have not added in your react app.

---

## 4. Terraform File structure must be
```bash
TERRAFORM-DIR/
├── main.tf                 # Defines AWS resources (S3, IAM, CodePipeline, CodeBuild)
├── variables.tf            # Input variables for the Terraform project
├── outputs.tf              # Optional outputs like URLs or resource names
├── terraform.tfvars        # Actual values for variables (keep this out of GitHub)
├── backend.tf              # (Optional) Remote state backend config like S3
├── buildspec.yml           
│
├── test/                  # Terratest files
│   └── main_test.go    # Go test file for testing infrastructure
│
├── scripts/                # (Optional) Shell scripts for automation
```
---

## 5. Provision AWS Infrastructure

```bash
terraform init
terraform plan
terraform validate #optional
terraform apply
```

---



## 6. CI/CD Pipeline Flow

1. **Source Stage** – GitHub repo
2. **Build Stage** – CodeBuild runs `buildspec.yml`
3. **Deploy Stage** – Deploys to S3 bucket (static hosting)

---

## 7. Running Terratest

Install And Run:

```bash
cd test
go mod init test
go get github.com/gruntwork-io/terratest/modules/terraform
go get github.com/stretchr/testify/assert
go test -v
```

---

## 8. Make sure Permission of Your Bucket
-- GO TO PERMISSION TAB 
- Block public access (bucket settings) should be UN-TICK
- You must hav a Bucket Policy
```bash
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "PublicRead",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::<YOUR BUCKET NAME>/*"
        }
    ]
}
```
-- GO TO PROPERTIES -> STATIC WEBSITE HOSTING 
-- CLICK ON THE URL
## 📄 License

MIT © Santosh Rout
