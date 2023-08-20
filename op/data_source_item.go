package op

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/winebarrel/terraform-provider-op/onepassword"
)

func dataSourceItem() *schema.Resource {
	return &schema.Resource{
		ReadContext: readItemByIDOrTitle,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"id", "title"},
			},
			"title": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"id", "title"},
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readItemByIDOrTitle(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*onepassword.Client)
	var idOrTitle string

	if _, ok := d.GetOk("id"); ok {
		idOrTitle = d.Get("id").(string)
	} else if _, ok := d.GetOk("title"); ok {
		idOrTitle = d.Get("title").(string)
	} else {
		return diag.Errorf("fatal: no id or title specified for item data source")
	}

	item, err := client.GetItem(idOrTitle)
	fields := map[string]any{}

	if err != nil {
		return diag.FromErr(err)
	}

	for _, f := range item.Fields {
		switch f.ID {
		case "username":
			d.Set("username", f.Value) //nolint:errcheck
		case "password":
			d.Set("password", f.Value) //nolint:errcheck
		default:
			if f.Label != nil {
				fields[*f.Label] = f.Value
			}
		}
	}

	d.SetId(item.ID)        //nolint:errcheck
	d.Set("fields", fields) //nolint:errcheck

	return nil
}
