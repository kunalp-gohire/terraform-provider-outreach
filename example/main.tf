terraform {
  required_providers{
      outreach ={
          version ="1.0"
          source = "outreach.com/edu/outreach"
      }
  }
}

provider "outreach" {
    
}

data "outreach_users" "user1"{
   email="kpgkunalgohire12@gmail.com"
}

output "users" {
  value = data.outreach_users.user1
}

resource "outreach_resource_user" "user"{
  email= "kpgkunalgohire@gmail.com"
  firstname="User11"
  lastname= "Test11"
  locked= true
}
resource "outreach_resource_user" "user1"{
  email= "kpgkunalgohire123@gmail.com"
  firstname="User123"
  lastname= "Test123"
  locked= true
}
output "user_instance1" {
  value = outreach_resource_user.user1
}


output "user_instance" {
  value = outreach_resource_user.user
}
