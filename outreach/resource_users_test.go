package outreach

import (
	"fmt"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "email", "asthayuno45@gmail.com"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "firstname", "Test"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "lastname", "User"),
				),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
resource "outreach_resource_user" "user1" {
	email= "asthayuno45@gmail.com"
  firstname="Test"
  lastname= "User"
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
						"outreach_resource_user.user2", "email", "kpgkunalgohire22@gmail.com"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "firstname", "UserU"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "lastname", "TestU"),
				),
			},
			{
				Config: testAccCheckUserUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserDataExists("outreach_resource_user.user2"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "email", "kpgkunalgohire22@gmail.com"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "firstname", "TestU"),
					resource.TestCheckResourceAttr(
						"outreach_resource_user.user2", "lastname", "UserU"),
				),
			},
		},
	})
}

func testAccCheckUserUpdatePre() string {
	return fmt.Sprintf(`
resource "outreach_resource_user" "user2" {
  email= "kpgkunalgohire22@gmail.com"
  firstname="UserU"
  lastname= "TestU"
  locked= true
	
}
`)
}

func testAccCheckUserUpdatePost() string {
	return fmt.Sprintf(`
resource "outreach_resource_user" "user2" {
  email= "kpgkunalgohire22@gmail.com"
  firstname="TestU"
  lastname= "UserU"
  locked= true
}
`)
}
