# outreach_resource_user ( Resource)
   The outreach_user  resource allows you to create and manage Outreach users.

## Example Usage
   ``` hcl
resource "outreach_user" "user" {
  email       = "test@example.com"
  firstname   = "User"
  lastname    = "Test"
  locked      = true
  title       = "Test"
}

output "user_instance" {
  value = outreach_user.user
}


   ```

## Schema

### Required
* `email`       (string)  - The email id associated with the user account.
* `firstname`   (string)  - First name of the User. 
* `lastname`    (string)  - Last Name / Family Name / Surname of the User. 
* `locked`      (boolean) - User account status.

### Optional
* `title`       (string) - Job title of user in organization.

### Read-Only 
* `id`          (int)   - Server generated user ID of user.
* `username` (string)   - The server generated username.




