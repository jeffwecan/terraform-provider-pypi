package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

	"github.com/jeffwecan/terraform-provider-pypi/pypi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pypi.Provider})
}
