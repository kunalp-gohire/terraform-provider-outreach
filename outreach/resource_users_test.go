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
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "email", "kpgkunalgohire2122@gmail.com"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "firstname", "Test2244"),
					resource.TestCheckResourceAttr("outreach_resource_user.user1", "lastname", "User2222"),
				),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
resource "outreach_resource_user" "user1" {
  email= "kpgkunalgohire2122@gmail.com"
  firstname="Test2244"
  lastname= "User2222"
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
