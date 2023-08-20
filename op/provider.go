package op

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/winebarrel/terraform-provider-op/onepassword"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"command": {
				Description: "`op` command path.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "op",
			},
		},
		ConfigureContextFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"op_item": dataSourceItem(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	cmd := d.Get("command").(string)
	client, err := onepassword.NewClient(cmd)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}
