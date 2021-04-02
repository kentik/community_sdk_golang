//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kentik/community_sdk_golang/terraform-provider-kentik-cloudexport/internal/provider"
)

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: provider.New}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/providers/kentik/kentik-cloudexport", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		plugin.Serve(opts)
	}
}
