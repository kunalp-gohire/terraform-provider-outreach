terraform {
  required_providers{
      outreach ={
          version ="1.0"
          source = "outreach.com/edu/outreach"
      }
  }
}

provider "outreach" {
    acc_token= "Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ2aWtyYW0rc2FuZGJveEBjbGV2ZXJ0YXAuY29tIiwiaWF0IjoxNjIxNDIwMDE2LCJleHAiOjE2MjE0MjcyMTYsImJlbnRvIjoiYXBwMWUiLCJvcmdfdXNlcl9pZCI6MSwiYXVkIjoiQ2xldmVyVGFwIC0gc2FuZGJveCIsInNjb3BlcyI6IkFQQT0iLCJvcmdfZ3VpZCI6IjljOThlNWU0LWFkOTItNDc0Yy1hNmQ5LTdmZDY0MmZjNTBkYSIsIm9yZ19zaG9ydG5hbWUiOiJjbGV2ZXJ0YXBzYW5kYm94MiJ9.lHowFTW6z4PdxJNiPxuDJsKIG94r4j3pZJ1bZJ1e6mo"
}
data "outreach_users" "all"{
   id=2
}

output "users" {
  value = data.outreach_users.all
}
