terraform {
  required_providers{
      outreach ={
          version ="1.0"
          source = "outreach.com/edu/outreach"
      }
  }
}

provider "outreach" {
    acc_token= "Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ2aWtyYW0rc2FuZGJveEBjbGV2ZXJ0YXAuY29tIiwiaWF0IjoxNjIxNTc0MDQyLCJleHAiOjE2MjE1ODEyNDIsImJlbnRvIjoiYXBwMWUiLCJvcmdfdXNlcl9pZCI6MSwiYXVkIjoiQ2xldmVyVGFwIC0gc2FuZGJveCIsInNjb3BlcyI6IkFQQT0iLCJvcmdfZ3VpZCI6IjljOThlNWU0LWFkOTItNDc0Yy1hNmQ5LTdmZDY0MmZjNTBkYSIsIm9yZ19zaG9ydG5hbWUiOiJjbGV2ZXJ0YXBzYW5kYm94MiJ9.w6Ibt5wYRHfu8eTHPL93NdWZ0GFg50eTC83URkusjfU"
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
