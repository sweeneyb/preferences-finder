package main

import (
	pf "github.com/sweeneyb/preferences-finder/terraform-provider/preferences_finder"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return pf.Provider()
		},
	})
}
