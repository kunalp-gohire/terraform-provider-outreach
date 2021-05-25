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

resource "outreach_resource_user" "import"{
}

output "user_import"{
    value=outreach_resource_user.import
}