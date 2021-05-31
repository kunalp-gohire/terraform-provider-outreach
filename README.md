# Outreach Terraform Provider

This terraform provider allows to perform Create ,Read ,Update, Import and Lock Outreach User(s). 


## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Outreach API Access](https://www.outreach.io/product/platform/api)
* Outreach Admin Account (Outreach Admin of organization can grant the Admin access to the user and can give API access credentials)


## Steps For Authorization Code Generation
 :heavy_exclamation_mark:  [IMPORTANT] : To test and run the Outreach Provider requires  the API access credentials<br><br>

1. <br>


## Initialise Outreach Provider in local machine 
1. Clone the repository  to $GOPATH/src/github.com/kunalp-gohire/Outreach_Terraform_Provider/tree/main<br>
2. Add the Client ID, Client Secret, Redirect URI and authorization code to respective fields in `main.tf` <br>
3. Run the following command :
 ```golang
go mod init terraform-provider-zoom
go mod tidy
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation
1. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
```
~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
``` 
Command: 
```bash
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/outreach/1.0/[OS_ARCH]
```
For eg. `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/outreach/1.0/windows_amd64`<br>

2. Run `go build -o terraform-provider-outreach.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-zoom.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\outreach\1.0\[OS_ARCH]
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


[OR]

1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Run the Terraform provider

#### Create User
1. Add the user email, first name, last name, locked, title and phone number in the respective field in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created and an user seat claim mail has been sent to the user.

#### Update the user
Update the data of the user in the `main.tf` file and apply using `terraform apply`

#### Read the User Data
Add data and output blocks in the `main.tf` file after that add email field  and user email  and run `terraform plan` to read user data

#### Lock/Unlock the user
Change the value of ‘locked’  from `false` to `true` or vice versa and run `terraform apply`.

#### Delete the user
 :heavy_exclamation_mark:  [IMPORTANT] : Outreach doesn’t provide an API to delete users, instead of delete they suggest to lock the user. To delete user from state file follow below instructions<br><br>
Delete the resource block of the particular user from the `main.tf` file and run `terraform apply`.

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
          version ="1.0"
          source = "outreach.com/edu/outreach"
      }
  }
}

provider "outreach" {
    
}

data "outreach_users" "user1"{
   email="[USER_EMAIL]"
}

output "user_data" {
  value = data.outreach_users.user1
}

resource "outreach_resource_user" "user"{
  email= "[USER_EMAIL]"
  firstname="[USER_FIRST_NAME]"
  lastname= "[USER_LAST_NAME]"
  locked= [USER_ACCOUNT_STATUS]
  phonenumber= "[USER_PHONE_NUMBER]"
  title="[USER_JOB_TITLE]"
}

output "user_instance" {
  value = outreach_resource_user.user
}

resource "outreach_resource_user" "user1"{
  email= "[USER_EMAIL]"
  firstname="[USER_FIRST_NAME]"
  lastname= "[USER_LAST_NAME]"
  locked= [USER_ACCOUNT_STATUS]
  phonenumber= "[USER_PHONE_NUMBER]"
  title="[USER_JOB_TITLE]"
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
