package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/demvsystems/terraform-provider-mongodbatlas/mongodbatlas"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mongodbatlas.Provider})
}
