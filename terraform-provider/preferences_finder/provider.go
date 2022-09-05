package preferences_finder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"preferences_finder_works": dataSourceWorks(),
			"pf_works":                 dataSourceWorks(),
		},
	}
}
