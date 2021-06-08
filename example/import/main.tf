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

resource "outreach_resource_user" "kunalgohire" {
  email     = "kunalp-gohire@gmail.com"
  firstname = "kunalp"
  lastname  = "Test12345"
  locked    = true
  title     = "Temp"
}


output "user_import" {
  value = outreach_resource_user.kunalgohire
}
