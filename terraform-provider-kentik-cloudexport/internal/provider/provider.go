package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

func New() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apiurl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_URL", nil),
				Description: "Custom apiserver url can be specified either by apiurl attribute or KTAPI_URL environment variable. If not specified, default of <https://cloudexports.api.kentik.com> will be used",
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_EMAIL", nil),
				Description: "Authorization. Either email attribute or KTAPI_AUTH_EMAIL environment variable is required",
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_TOKEN", nil),
				Description: "Authorization. Either token attribute or KTAPI_AUTH_TOKEN  environment variable is required",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"kentik-cloudexport_list": dataSourceCloudExportList(),
			"kentik-cloudexport_item": dataSourceCloudExportItem(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"kentik-cloudexport_item": resourceCloudExport(),
		},
	}

	p.ConfigureContextFunc = configure

	return p
}

func configure(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	email := d.Get("email").(string)
	token := d.Get("token").(string)
	apiurl, apiurlOK := d.GetOk("apiurl")

	if !apiurlOK {
		return newClient(email, token, ""), nil
	}
	return newClient(email, token, apiurl.(string)), nil
}

func newClient(email, token, url string) *kentikapi.Client {
	cfg := kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
		APIURL:    url,
	}
	return kentikapi.NewClient(cfg)
}
