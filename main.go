package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/winebarrel/terraform-provider-op/op"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name redash

func main() {
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: op.Provider,
		ProviderAddr: "registry.terraform.io/winebarrel/op",
		Debug:        *debug,
	}

	plugin.Serve(opts)
}
