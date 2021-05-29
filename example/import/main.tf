terraform {
  required_providers {
    outreach = {
      version = "1.0"
      source  = "outreach.com/edu/outreach"
    }
  }
}

provider "outreach" {

}

resource "outreach_resource_user" "import" {
   
}

output "user_import" {
  value = outreach_resource_user.import
}
