package outreach

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUser_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserDataExists("outreach_resource_user.user1"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "email", "kpgkunalgohire2222@gmail.com"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "firstname", "Test12244"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "lastname", "User12222"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "locked", "true"),
				),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
resource "outreach_resource_user" "user1" {
  email= "kpgkunalgohire2222@gmail.com"
  firstname="Test12244"
  lastname= "User12222"
  locked= true
}
`)
}

func TestAccUser_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserDataExists("outreach_resource_user.user2"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "email", "kpgkunal44@gmail.com"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "firstname", "User444"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "lastname", "TestU244"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "locked", "true"),
				),
			},
			{
				Config: testAccCheckUserUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserDataExists("outreach_resource_user.user2"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "email", "kpgkunal44@gmail.com"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "firstname", "TESTU21"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "lastname", "USERU21"),
				),
			},
		},
	})
}

func testAccCheckUserUpdatePre() string {
	return fmt.Sprintf(`
resource "outreach_resource_user" "user2" {
  email= "kpgkunal44@gmail.com"
  firstname="User444"
  lastname= "TestU244"
  locked= true	
}
`)
}

func testAccCheckUserUpdatePost() string {
	return fmt.Sprintf(`
resource "outreach_resource_user" "user2" {
  email= "kpgkunal44@gmail.com"
  firstname="TESTU21"
  lastname= "USERU21"
  locked= true
}
`)
}

