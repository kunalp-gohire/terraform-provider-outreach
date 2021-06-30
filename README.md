* This terraform provider allows to perform Create ,Read ,Update, Import and Lock    Outreach User(s). 
* To fetch and import a user server generated User ID is needed.
* Outreach doesn't provide an API to delete  a user.


## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Outreach API Access](https://www.outreach.io/product/platform/api)
* Outreach Admin Account (Outreach Admin of organization can grant the Admin access to the user and can give API access credentials)
* [Outreach API documentation](https://api.outreach.io/api/v2/docs#getting-started)

## Application Setup 

*  Request the Outreach API access credentials from organizations Outreach admin
(client secret, client ID, redirect URI) and account having admin access.
*  Access token is required to make API requests. 
*  To automate access token generation refresh token is required. 

### Steps to Generate the Refresh Token

1. Replace the  `<Application_Identifier>`  with client ID/application ID and `<Application_Redirect_URI>` with redirect URL in the following URL.
 ```
https://api.outreach.io/oauth/authorize?client_id=<Application_Identifier>&redirect_uri=<Application_Redirect_URI>&response_type=code&scope=users.all
 ```
2. Request an authorization code from an Outreach customer by redirecting them to the       above URL.
3. Once the Outreach customer has clicked Authorize they will be redirected back to       your link’s redirect_uri with a code(authorization code) query parameter.
4. Use that authorization code to make ‘POST’ request using the [Postman](https://www.postman.com/) 
5. To receive a refresh token, make a POST request to the following URL with the following parameters:
```curl
curl https://api.outreach.io/oauth/token \
  -X POST \
  -d client_id=<Application_Identifier> \
  -d client_secret=<Application_Secret> \
  -d redirect_uri=<Application_Redirect_URI> \
  -d grant_type=authorization_code \
  -d code=<Authorization_Code>
```
6. Refer the official [Outreach API documentation](https://api.outreach.io/api/v2/docs#authentication) 




## Building The Provider 
1. Clone the ‘main’ branch of the [Terraform Outreach Provider](https://github.com/kunalp-gohire/terraform-provider-outreach.git)<br>
```
https://github.com/kunalp-gohire/terraform-provider-outreach.git
```

2. Run the following commands :
 ```golang
go mod init terraform-provider-outreach
go mod tidy
```
3. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Managing plugins for terraform
1. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
```bash
mkdir -p ~/.terraform.d/plugins/outreach.com/edu/outreach/1.0/windows_amd64
```
2. Run `go build -o terraform-provider-outreach.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-outreach.exe %APPDATA%\terraform.d\plugins\outreach.com\edu\outreach\1.0\windows_amd64
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


[OR]

1. Download required binaries <br>
2. move binary `%APPDATA%\terraform.d\plugins\outreach.com\edu\outreach\1.0\windows_amd64`


## Working with terraform


### Authentication
1.  Add `terraform` block and `provider` block as shown in [example usage](#example-usage).
2. Add the Client ID, Client Secret, Redirect URI and Refresh Token to respective fields in provider block<br>

[OR]

2. Set the the global environment variables as Follow:
   * Client ID      - OUTREACH_CLIENT_ID
   * Client Secrete - OUTREACH_CLIENT_SECRET
   * Refresh Token  - OUTREACH_REFRESH_TOKEN  
   * Redirect URL   - OUTREACH_REDIRECT_URI 


### Basic Terraform Commands
* `terraform init`     - Prepare your working directory for other commands
* `terraform validate` - Check whether the configuration is valid
* `terraform plan`     - Show changes required by the current configuration
* `terraform apply`    - Create or update infrastructure
* `terraform destroy`  - Destroy previously-created infrastructure

### Create User
1. Add the user email, first name, last name, locked, title and phone number in the respective field in resource block as shown in [example usage](#example-usage).
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created and an user seat claim mail has been sent to the user.

### Update the user
1. Update the data of the user in the resource block and apply using `terraform apply`.
2. First name, last name, locked status, and title can be updated.
3. Change the value of 'locked'  from `false` to `true` or vice versa in the resource block and run `terraform apply`.
4. Email can't be updated through API. 

### Read the User Data
1. Add data and output blocks as shown in [example usage](#example-usage) after that add id field  and user ID  and run `terraform plan` to read user data.


### Delete the user
*Outreach doesn’t provide an API to delete users, instead of delete they suggest to lock the user.*<br><br>
To delete user from state file follow below instructions:<br><br>
1. Delete or comment the resource block of the particular user and run `terraform apply`.

### Import a User Data
1. Write manually a resource configuration block as shown in [example usage](#example-usage), to which the imported object will be mapped or define the empty resource block.
2. Run the command `terraform outreach_resource_user.import  [user ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in the resource block.
```
terraform outreach_resource_user.import  [user ID]
```




## Example Usage<a id="example-usage"></a>
```terraform
terraform {
  required_providers{
      outreach ={
          version = "1.0"
          source  = "outreach.com/edu/outreach"
      }
  }
}

provider "outreach" {
    
}

data "outreach_user" "user1"{
   id = 5
}

output "user_data" {
  value = data.outreach_user.user1
}

resource "outreach_user" "user"{
  email       = "user@example.com"
  firstname   = "Test"
  lastname    = "User"
  locked      = false
  title       = "Developer"
}

output "user_instance" {
  value = outreach_user.user
}

```

## Argument Reference

* `email`       (required, string) - The email id associated with the user account. Required for user creation.
* `firstname`   (required, string) - First name of the User. Required for user creation.
* `lastname`    (required, string) - Last Name / Family Name / Surname of the User. Required for user creation.
* `locked`      (boolean)- User account status. Required for user creation. Default value is false. 
                           If it is not provided at the time of user creation, then the created user will be active. 
* `phonenumber` (string) - Phone number of user. 
* `title`       (string) - Job title of user in organization. 
* `id`          (int)    - Server generated user ID. Required for fetch the user data using data block 
                           and import a user.                               
* `client_id`     (required, string) - Outreach Client ID/ Application ID. It can be set as environment 
                                       variable `OUTREACH_CLIENT_ID`. 
* `client_secret` (required, string) - Outreach Client secret ID. It can be set as environment 
                                       variable `OUTREACH_CLIENT_SECRET`.
* `redirect_url`  (required, string) - Outreach Application redirect URL. It can be set as environment 
                                       variable `OUTREACH_REDIRECT_URI`.
* `refresh_token` (required, string) - Refresh token. It can be set as environment variable 
                                       `OUTREACH_REFRESH_TOKEN`.




