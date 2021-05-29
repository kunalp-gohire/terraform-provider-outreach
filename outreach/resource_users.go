package outreach

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-outreach/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
				Type:     schema.TypeBool,
				Optional: true,
			},
			"uid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				email := d.Id()
				c := m.(*client.Client)
				user, err := c.GetDataSourceUser(email)
				if err != nil {
					 return nil, fmt.Errorf("%v ",err)
				}
				d.Set("email", user.Attributes.Email)
				d.Set("firstname", user.Attributes.FirstName)
				d.Set("lastname", user.Attributes.LastName)
				d.Set("locked", user.Attributes.Locked)
				d.Set("username", user.Attributes.UserName)
				d.Set("id", user.ID)
				UserId := user.ID
				uid := fmt.Sprintf("%v", UserId)
				d.SetId(uid)
				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	req_json := client.Data{
		Data: client.User{
			Type: "user",
			Attributes: client.Attributes{
				Email:     d.Get("email").(string),
				FirstName: d.Get("firstname").(string),
				LastName:  d.Get("lastname").(string),
				Locked:    d.Get("locked").(bool),
			},
		},
	}
	user, err := c.CreateUser(req_json)
	if err != nil {
		return diag.FromErr(err)
	}
	id := strconv.Itoa(user.Data.ID)
	d.Set("email", user.Data.Attributes.Email)
	d.Set("firstname", user.Data.Attributes.FirstName)
	d.Set("lastname", user.Data.Attributes.LastName)
	d.Set("locked", user.Data.Attributes.Locked)
	d.Set("username", user.Data.Attributes.UserName)
	d.SetId(id)
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	id := d.Id()
	user, err := c.GetUserData(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("uid", user.Data.ID)
	d.Set("email", user.Data.Attributes.Email)
	d.Set("firstname", user.Data.Attributes.FirstName)
	d.Set("lastname", user.Data.Attributes.LastName)
	d.Set("locked", user.Data.Attributes.Locked)
	d.Set("username", user.Data.Attributes.UserName)
	d.SetId(strconv.Itoa(user.Data.ID))
	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Can't update email",
			Detail:   "Can't update email",
		})
		return diags
	}
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if d.HasChange("lastname") || d.HasChange("firstname") || d.HasChange("locked") {
		req_json := client.Data{
			Data: client.User{
				Type: "user",
				ID:   id,
				Attributes: client.Attributes{
					Email:     d.Get("email").(string),
					FirstName: d.Get("firstname").(string),
					LastName:  d.Get("lastname").(string),
					Locked:    d.Get("locked").(bool),
				},
			},
		}
		UserID := d.Id()
		d.SetId(UserID)
		c.UpdateUser(UserID, req_json)
	}
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
