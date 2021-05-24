package outreach

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceContact() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// "attributes": {
			// 	Type:     schema.TypeSet,
			// 	Computed: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
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
			// "createdat":{
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			"locked": {
				Type:   schema.TypeBool,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// "title": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			// "updatedat":{
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			// },
			// },
			// },
		},
		ReadContext: dataSourceUserRead,
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)
	UserId := d.Get("id")
	uid:=fmt.Sprintf("%v" ,UserId)

	user, err := c.GetUserData(uid)
	if err != nil {
		return diag.FromErr(err)
	}
	
	// uis := make([]interface{}, 1)
	// ui := make(map[string]interface{})
	// ui["email"] = user.Data.Attributes.Email
	// ui["first_name"] = user.Data.Attributes.FirstName
	// ui["last_name"] = user.Data.Attributes.LastName
	// // ui["createdat"]=user.Data.Attributes.CreateAt
	// ui["locked"] = user.Data.Attributes.Locked
	// ui["username"] = user.Data.Attributes.UserName
	// ui["title"] = user.Data.Attributes.Title
	// // ui["updatedat"]=user.Data.Attributes.UpdatedAt
	// uis[0] = ui
	// d.Set("attributes", uis)

	d.Set("email",user.Data.Attributes.Email)
	d.Set("firstname",user.Data.Attributes.FirstName)
	d.Set("lastname",user.Data.Attributes.LastName)
	d.Set("locked",user.Data.Attributes.Locked)
	d.Set("username",user.Data.Attributes.UserName)
	// d.Set("title",user.Data.Attributes.Title)
	
	d.SetId(uid)
	return diags
}
