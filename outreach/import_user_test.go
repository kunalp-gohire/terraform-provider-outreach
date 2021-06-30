package outreach

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUser_import_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserImporterBasic(),
			},
			{
				ResourceName:      "outreach_user.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckUserImporterBasic() string {
	return fmt.Sprintf(`
	resource "outreach_user" "import" {
		email= "kpgkunalgohire1888@gmail.com"
		firstname="Test1888"
		lastname= "User1888"
		locked= true
	}
	`)
}
