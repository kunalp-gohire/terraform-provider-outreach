package outreach

import (
	"context"
	"fmt"

	// "fmt"
	// "encoding/json"
	// "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	// "net/http"
	// "os"
	// "strconv"
	// "time"
)

func dataSourceContact() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firstname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lastname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: dataSourceUserRead,
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)
	UserId := d.Get("id")
	uid := fmt.Sprintf("%v", UserId)

	user, err := c.GetUserData(uid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("email", user.Data.Attributes.Email)
	d.Set("firstname", user.Data.Attributes.FirstName)
	d.Set("lastname", user.Data.Attributes.LastName)
	d.Set("locked", user.Data.Attributes.Locked)
	d.Set("username", user.Data.Attributes.UserName)
	d.SetId(uid)
	return diags
}
