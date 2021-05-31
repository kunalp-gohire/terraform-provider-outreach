package outreach

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-outreach/client"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
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
			"phonenumber": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: dataSourceUserRead,
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	email := d.Get("email")
	useremail := fmt.Sprintf("%v", email)
	user, err := c.GetDataSourceUser(useremail)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("email", user.Attributes.Email)
	d.Set("firstname", user.Attributes.FirstName)
	d.Set("lastname", user.Attributes.LastName)
	d.Set("locked", user.Attributes.Locked)
	d.Set("username", user.Attributes.UserName)
	d.Set("id", user.ID)
	d.Set("phonenumber",user.Attributes.PhoneNumber)
	d.Set("title",user.Attributes.Title)
	UserId := user.ID
	uid := fmt.Sprintf("%v", UserId)
	d.SetId(uid)
	return diags
}
