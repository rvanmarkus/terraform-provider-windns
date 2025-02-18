package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/nrkno/terraform-provider-windns/internal/provider"
)

// these will be set by the goreleaser configuration
// to appropriate values for the compiled binary.
var version = "dev"

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()
	opts := &plugin.ServeOpts{
		Debug: debugMode,

		ProviderAddr: "registry.terraform.io/nrkno/windns",
		ProviderFunc: provider.Provider(version),
	}

	plugin.Serve(opts)
}
