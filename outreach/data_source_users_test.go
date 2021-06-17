package outreach

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckUserDataExists("data.outreach_user.test"),
					resource.TestCheckResourceAttr(
						"data.outreach_user.test", "email", "kpgkunalgohire123@gmail.com"),
					resource.TestCheckResourceAttr(
						"data.outreach_user.test", "firstname", "User123"),
					resource.TestCheckResourceAttr(
						"data.outreach_user.test", "lastname", "Test123"),
					resource.TestCheckResourceAttr(
						"data.outreach_user.test", "locked", "true"),
					resource.TestCheckResourceAttr(
						"data.outreach_user.test", "username", "User123_Test123"),
				),
			},
		},
	})
}

func testAccCheckUserDataExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		return nil
	}
}

func testAccUserDataSourceConfig() string {
	return fmt.Sprintf(`
data "outreach_user" "test" {
	id = 5
}
`)
}
