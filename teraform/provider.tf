terraform {
  required_providers {
    vultr = {
      source = "vultr/vultr"
      version = "2.19.0"
    }
  }
}

provider "vultr" {
  api_key = var.api_key
  rate_limit = 100
  retry_limit = 3
}

 terraform {
   backend "s3" {
     bucket                      = "terraform-state-wosp"
     key                         = "terraform.tfstate"
     endpoint                    = "ewr1.vultrobjects.com"
     region                      = "us-east-1"
     skip_credentials_validation = true
   }
 }