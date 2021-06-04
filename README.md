# Outreach Terraform Provider

* This terraform provider allows to perform Create ,Read ,Update, Import and Lock    Outreach User(s). 
* To fetch and import a user server generated User ID is needed.
* Outreach doesn't provide an API to delete  a user.
* [Outreach API documentation](https://api.outreach.io/api/v2/docs#getting-started)

## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Outreach API Access](https://www.outreach.io/product/platform/api)
* Outreach Admin Account (Outreach Admin of organization can grant the Admin access to the user and can give API access credentials)

## Initialise Outreach Provider in local machine 
1. Clone the ‘main’ branch of the [Terraform Outreach Provider](https://github.com/kunalp-gohire/Outreach_Terraform_Provider.git)<br>
```
https://github.com/kunalp-gohire/Outreach_Terraform_Provider.git
```

2. Run the following commands :
 ```golang
go mod init terraform-provider-zoom
go mod tidy
```
3. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation
1. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
```
~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
``` 
```bash
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/outreach/1.0/[OS_ARCH]
```
For eg. `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/outreach/1.0/windows_amd64`<br>

2. Run `go build -o terraform-provider-outreach.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-outreach.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\outreach\1.0\[OS_ARCH]
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


[OR]

1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Run the Terraform provider

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

### Authentication
1. Add the Client ID, Client Secret, Redirect URI and Refresh Token to respective fields in `main.tf` <br>

[OR]

1. Set the the global environment variables as Follow:
   * Client ID - outreach_client_id
   * Client Secrete - outreach_client_secrete
   * Refresh Token - outreach_refresh_token  

#### Create User
1. Add the user email, first name, last name, locked, title and phone number in the respective field in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created and an user seat claim mail has been sent to the user.

#### Update the user
1. Update the data of the user in the `main.tf` file and apply using `terraform apply`

#### Read the User Data
1. Add data and output blocks in the `main.tf` file after that add email field  and user email  and run `terraform plan` to read user data

#### Lock/Unlock the user
1. Change the value of ‘locked’  from `false` to `true` or vice versa and run `terraform apply`.

#### Delete the user
 :heavy_exclamation_mark:  [IMPORTANT] : Outreach doesn’t provide an API to delete users, instead of delete they suggest to lock the user. To delete user from state file follow below instructions:<br><br>
1. Delete or comment the resource block of the particular user from the `main.tf` file and run `terraform apply`.

#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped or define the empty resource block.
2. Run the command `terraform outreach_resource_user.import “[EMAIL_ID]”`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in the resource block.


### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` for unit testing and for acceptance testing run command `TF_ACC=1 go test` . These commands will give combined test results for the execution or errors if any failure occurs.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`


## Example Usage
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

data "outreach_users" "user1"{
   email = "[USER_EMAIL]"
}

output "user_data" {
  value = data.outreach_users.user1
}

resource "outreach_resource_user" "user"{
  email               = "[USER_EMAIL]"
  firstname         = "[USER_FIRST_NAME]"
  lastname         = "[USER_LAST_NAME]"
  locked             = [USER_ACCOUNT_STATUS]
  phonenumber = "[USER_PHONE_NUMBER]"
  title                  = "[USER_JOB_TITLE]"
}

output "user_instance" {
  value = outreach_resource_user.user
}

resource "outreach_resource_user" "user1"{
  email               = "[USER_EMAIL]"
  firstname         = "[USER_FIRST_NAME]"
  lastname         = "[USER_LAST_NAME]"
  locked             = [USER_ACCOUNT_STATUS]
  phonenumber = "[USER_PHONE_NUMBER]"
  title                  = "[USER_JOB_TITLE]"

}

output "user_instance1" {
  value = outreach_resource_user.user1
}
```

## Argument Reference

* `email` (string)     - The email id associated with the user account.
* `firstname` (string) - First name of the User.
* `lastname` (string) - Last Name / Family Name / Surname of the User.
* `locked` (boolean) - User account status.
* `phonenumber` (string) - Phone number of user.
* `title` (string) -  Job title of user in organization. 

