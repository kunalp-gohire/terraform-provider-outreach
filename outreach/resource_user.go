package outreach

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"terraform-provider-outreach/client"
	"time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateEmail(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value := v.(string)
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !(emailRegex.MatchString(value)) {
		errs = append(errs, fmt.Errorf("expected emailId is not valid  %s", k))
		return warns, errs
	}
	return
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"email": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateEmail,
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
			"phonenumber": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
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
			StateContext: resourceUserImporter,
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
				Title:     d.Get("title").(string),
			},
		},
	}
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		user, err := c.CreateUser(req_json)
		if err != nil {
			if c.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		id := strconv.Itoa(user.Data.ID)
		d.Set("email", user.Data.Attributes.Email)
		d.Set("firstname", user.Data.Attributes.FirstName)
		d.Set("lastname", user.Data.Attributes.LastName)
		d.Set("locked", user.Data.Attributes.Locked)
		d.Set("username", user.Data.Attributes.UserName)
		d.Set("title", user.Data.Attributes.Title)
		d.Set("phonenumber", user.Data.Attributes.PhoneNumber)
		d.SetId(id)
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	id := d.Id()
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		user, err := c.GetUserData(id)
		if err != nil {
			if c.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		d.Set("uid", user.Data.ID)
		d.Set("email", user.Data.Attributes.Email)
		d.Set("firstname", user.Data.Attributes.FirstName)
		d.Set("lastname", user.Data.Attributes.LastName)
		d.Set("locked", user.Data.Attributes.Locked)
		d.Set("username", user.Data.Attributes.UserName)
		d.Set("title", user.Data.Attributes.Title)
		d.Set("phonenumber", user.Data.Attributes.PhoneNumber)
		d.SetId(strconv.Itoa(user.Data.ID))
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		if strings.Contains(retryErr.Error(), "User Does Not Exist , StatusCode = 404") {
			d.SetId("")
			return diags
		}
		return diag.FromErr(retryErr)
	}
	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Can't update email",
			Detail:   "Can't update email through API",
		})
		return diags
	}
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if d.HasChange("lastname") || d.HasChange("firstname") || d.HasChange("locked") || d.HasChange("phonenumber") || d.HasChange("title") {
		req_json := client.Data{
			Data: client.User{
				Type: "user",
				ID:   id,
				Attributes: client.Attributes{
					Email:     d.Get("email").(string),
					FirstName: d.Get("firstname").(string),
					LastName:  d.Get("lastname").(string),
					Locked:    d.Get("locked").(bool),
					Title:     d.Get("title").(string),
				},
			},
		}
		retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
			UserID := d.Id()
			d.SetId(UserID)
			if _, err = c.UpdateUser(UserID, req_json); err != nil {
				if c.IsRetry(err) {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if retryErr != nil {
			time.Sleep(2 * time.Second)
			return diag.FromErr(retryErr)
		}
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func resourceUserImporter(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()
	uid := fmt.Sprintf("%v", id)
	c := m.(*client.Client)
	user, err := c.GetUserData(uid)
	if err != nil {
		return nil, err
	}
	d.Set("uid", user.Data.ID)
	d.Set("email", user.Data.Attributes.Email)
	d.Set("firstname", user.Data.Attributes.FirstName)
	d.Set("lastname", user.Data.Attributes.LastName)
	d.Set("locked", user.Data.Attributes.Locked)
	d.Set("username", user.Data.Attributes.UserName)
	d.Set("title", user.Data.Attributes.Title)
	d.Set("phonenumber", user.Data.Attributes.PhoneNumber)
	d.SetId(uid)
	return []*schema.ResourceData{d}, nil
}
