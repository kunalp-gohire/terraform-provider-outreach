package outreach

import (
	"context"
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
				Type:     schema.TypeString,
				Required: true,
			},
			
			"attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"createdAt":{
							Type:     schema.TypeString,
							Computed: true,
						},
						"locked":{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"username":{
							Type:     schema.TypeString,
							Computed: true,
						},
						"title":{
							Type:     schema.TypeString,
							Computed: true,
						},
						"updatedAt":{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		ReadContext: dataSourceUserRead,
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	
	return diags
}
