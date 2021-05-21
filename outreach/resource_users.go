package outreach

import (
	"context"
	"strconv"
	// "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"firstname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lastname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"locked": {
				Type:   schema.TypeBool,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		// UpdateContext: resourceUserUpdate,
		// DeleteContext: resourceUserDelete,
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	req_json := Data{
		Data:User{
             Type: "user",
			 Attributes: Attributes{
				 Email: d.Get("email").(string),
				 FirstName: d.Get("firstname").(string),
				 LastName: d.Get("lastname").(string),
			 },
		},
	}
	user, err := c.CreateUser(req_json)
	if err != nil {
		return diag.FromErr(err)
	}
	id:=user.Data.ID
	d.Set("email",user.Data.Attributes.Email)
	d.Set("firstname",user.Data.Attributes.FirstName)
	d.Set("lastname",user.Data.Attributes.LastName)
	d.Set("locked",user.Data.Attributes.Locked)
    d.Set("username",user.Data.Attributes.UserName)
	d.Set("id",user.Data.ID)
	d.SetId(strconv.Itoa(id))
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)
	id:=d.Id()
	user, err := c.GetUserData(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("email",user.Data.Attributes.Email)
	d.Set("firstname",user.Data.Attributes.FirstName)
	d.Set("lastname",user.Data.Attributes.LastName)
	d.Set("locked",user.Data.Attributes.Locked)
	d.Set("username",user.Data.Attributes.UserName)
	d.SetId(strconv.Itoa(user.Data.ID))
	return diags
}


