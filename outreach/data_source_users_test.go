package outreach

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
					testAccCheckUserDataExists("data.outreach_users.test"),
					resource.TestCheckResourceAttr(
						"data.outreach_users.test", "email", "kpgkunalgohire123@gmail.com"),
					resource.TestCheckResourceAttr(
						"data.outreach_users.test", "firstname", "User123"),
					resource.TestCheckResourceAttr(
						"data.outreach_users.test", "lastname", "Test123"),
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
	fmt.Println("create data block")
	return fmt.Sprintf(`
data "outreach_users" "test" {
	   email     = "kpgkunalgohire123@gmail.com"
}
`)
}
