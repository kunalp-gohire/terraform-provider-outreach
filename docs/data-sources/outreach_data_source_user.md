# outreach_user Data Source
   The outreach_user data source provides information about an existing Outreach user.

## Example Usage
   ``` hcl
data "outreach_user" "user" {
  id = 9
}

output "user_data" {
  value = data.outreach_user.user
}

   ```

## Schema

### Required
* `id` (int) - Server generated user ID of user.

### Read-Only 
* `email`       (string) - The email id associated with the user account.
* `firstname`   (string) - First name of the User. 
* `lastname`    (string) - Last Name / Family Name / Surname of the User. 
* `locked`      (boolean)- User account status.
* `phonenumber` (string) - Phone number of user. 
* `title`       (string) - Job title of user in organization.

