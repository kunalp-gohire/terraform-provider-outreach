package outreach

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-outreach/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTREACH_CLIENT_ID", nil),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTREACH_CLIENT_SECRET", nil),
			},
			"refresh_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTREACH_REFRESH_TOKEN", nil),
			},
			"redirect_uri": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTREACH_REDIRECT_URI", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"outreach_user": resourceUser(),
		},
		DataSourcesMap:       map[string]*schema.Resource{"outreach_user": dataSourceUsers()},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client_id := d.Get("client_id").(string)
	client_secret := d.Get("client_secret").(string)
	refresh_token := d.Get("refresh_token").(string)
	redirect_uri := d.Get("redirect_uri").(string)
	c, err := client.NewClient(client_id, client_secret, refresh_token, redirect_uri)
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
