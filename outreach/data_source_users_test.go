package outreach

import (
	"fmt"
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
					resource.TestCheckResourceAttr(
						"data.outreach_users.test", "email", "kpgkunalgohire123@gmail.com"),
				),
				
			},
		},
	})
}

func testAccUserDataSourceConfig() string {
	fmt.Println("create data block")
	return fmt.Sprintf(`
data "outreach_users" "test" {
	   email     = "kpgkunalgohire123@gmail.com"
}
`)
}
