variable "region" {
  default = "ap-south-1"
}

variable "artifact_bucket" {
  description = "S3 bucket name for storing pipeline artifacts"
  type        = string
}

variable "github_owner" {
  description = "GitHub repository owner"
  type        = string
}

variable "github_repo" {
  description = "GitHub repository name"
  type        = string
}

variable "github_branch" {
  description = "Branch to deploy from"
  type        = string
  default     = "main"
}

variable "github_token" {
  description = "GitHub OAuth token"
  type        = string
  sensitive   = true
}
