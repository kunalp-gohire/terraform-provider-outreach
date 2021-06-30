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
					testAccCheckUserDataExists("outreach_user.user1"),
					resource.TestCheckResourceAttr("outreach_user.user1", "email", "kpgkunalgohire88@gmail.com"),
					resource.TestCheckResourceAttr("outreach_user.user1", "firstname", "Test88"),
					resource.TestCheckResourceAttr("outreach_user.user1", "lastname", "User88"),
					resource.TestCheckResourceAttr("outreach_user.user1", "locked", "true"),
				),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
resource "outreach_user" "user1" {
  email= "kpgkunalgohire88@gmail.com"
  firstname="Test88"
  lastname= "User88"
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
					testAccCheckUserDataExists("outreach_user.user2"),
					resource.TestCheckResourceAttr(
						"outreach_user.user2", "email", "kpgkunalgohire99@gmail.com"),
					resource.TestCheckResourceAttr(
						"outreach_user.user2", "firstname", "User99"),
					resource.TestCheckResourceAttr(
						"outreach_user.user2", "lastname", "Test99"),
					resource.TestCheckResourceAttr(
						"outreach_user.user2", "locked", "true"),
				),
			},
			{
				Config: testAccCheckUserUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserDataExists("outreach_user.user2"),
					resource.TestCheckResourceAttr(
						"outreach_user.user2", "email", "kpgkunalgohire99@gmail.com"),
					resource.TestCheckResourceAttr(
						"outreach_user.user2", "firstname", "USER99"),
					resource.TestCheckResourceAttr(
						"outreach_user.user2", "lastname", "TEST99"),
				),
			},
		},
	})
}

func testAccCheckUserUpdatePre() string {
	return fmt.Sprintf(`
resource "outreach_user" "user2" {
  email= "kpgkunalgohire99@gmail.com"
  firstname="User99"
  lastname= "Test99"
  locked= true	
}
`)
}

func testAccCheckUserUpdatePost() string {
	return fmt.Sprintf(`
resource "outreach_user" "user2" {
   email= "kpgkunalgohire99@gmail.com"
   firstname="USER99"
   lastname= "TEST99"
   locked= true
}
`)
}
