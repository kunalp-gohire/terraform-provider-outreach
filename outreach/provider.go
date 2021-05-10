package outreach

import (
	"context"
	
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{


		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("client_id", nil),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("client_secret", nil),
			},
			"redirect_uri":&schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("redirect_uri", nil),
			},
		},
		// resouces block
		ResourcesMap: map[string]*schema.Resource{
			
		},

		// data block
		DataSourcesMap: map[string]*schema.Resource{
			
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// type Client struct{
// 	authToken  string
// }
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics){
    
	var diags diag.Diagnostics
	
	return _, diags
}