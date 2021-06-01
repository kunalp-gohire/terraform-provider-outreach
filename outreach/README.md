# Acceptance Testing

<https://www.terraform.io/docs/extend/testing/acceptance-tests/index.html>

#### BASIC COMMANDS TO RUN ACCEPTANCE TESTING

1-TF_ACC=1 go test -v <br/>
2-TF_ACC=1 go test -cover -v (to get idea about how much percentage of your code is tested) <br />


### DIFFERENT TESTING FUNCTION

<strong>1. testAccCheckUserDataExists </strong>

Checks the resource block or data block exist or not in terraform state file.<br />

<strong>2. TestAccUserDataSource_basic </strong>

Creates the data block and verifies that the returned resource attributes match.<br />

<strong>4. TestAccUser_Basic </strong>

Creates the resource block and verifies that the returned resource attributes match. <br />

<strong>5. TestAccUser_Update</strong>

Creates the resource block, updates the resource block attributes and verifies that the returned resource attributes matches. <br />
