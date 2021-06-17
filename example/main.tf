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

data "outreach_user" "user1" {
  id = 11
}

output "user_data" {
  value = data.outreach_user.user1
}

resource "outreach_user" "user" {
  email       = "kpgkunalgohire@gmail.com"
  firstname   = "User11"
  lastname    = "Test11"
  locked      = true
  title       = "Test"
}

output "user_instance" {
  value = outreach_user.user
}

resource "outreach_user" "user1" {
  email       = "kpgkunalgohire123@gmail.com"
  firstname   = "User123"
  lastname    = "Test123"
  locked      = true
  title       = "Test"
}

output "user_instance1" {
  value = outreach_user.user1
}

resource "outreach_user" "user3" {
  email     = "ashwini.gaddagi@clevertap.com"
  firstname = "Ashwini"
  lastname  = "Test12345"
  locked    = true
}

output "user_instance3" {
  value = outreach_user.user3
}
