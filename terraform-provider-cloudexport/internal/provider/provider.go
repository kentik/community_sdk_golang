package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"apiurl": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("KTAPI_URL", nil),
				},
				"email": &schema.Schema{
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_EMAIL", nil),
				},
				"token": &schema.Schema{
					Type:        schema.TypeString,
					Sensitive:   true,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_TOKEN", nil),
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

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		email := d.Get("email").(string)
		token := d.Get("token").(string)
		apiurl, apiurlOK := d.GetOk("apiurl")

		if !apiurlOK {
			return newClient(email, token, ""), nil
		}
		return newClient(email, token, apiurl.(string)), nil
	}
}

func newClient(email, token, url string) *kentikapi.Client {
	cfg := kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,
		APIURL:    url,
	}
	return kentikapi.NewClient(cfg)
}
