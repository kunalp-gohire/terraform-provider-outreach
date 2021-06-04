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
  email       = "kpgkunalgohire@gmail.com"
  firstname   = "User11"
  lastname    = "Test11"
  locked      = true
  phonenumber = ""
  title       = "Test"
}


output "user_import" {
  value = outreach_resource_user.import
}
