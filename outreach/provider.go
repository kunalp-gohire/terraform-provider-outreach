package outreach

import (
	"context"
    "terraform-provider-outreach/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("outreach_client_id", nil),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("outreach_client_secrete", nil),
			},
			"refresh_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("outreach_refresh_token", nil),
			},
			"acc_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("outreach_acc_token", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"outreach_resource_user":  resourceUser(),
		},
		DataSourcesMap:       map[string]*schema.Resource{"outreach_users": dataSourceUsers()},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client_id := d.Get("client_id").(string)
	client_secret := d.Get("client_secret").(string)
	refresh_token := d.Get("refresh_token").(string)
	acc_token := d.Get("acc_token").(string)
	c, err := client.NewClient(client_id, client_secret, refresh_token,acc_token)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Outreach client",
			Detail:   "Unable to authenticate user for authenticated Outreach client",
		})
		return nil, diags
	}
	return c, diags
}
