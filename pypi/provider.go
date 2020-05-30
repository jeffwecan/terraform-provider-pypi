package pypi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const DefaultPyPiAddress = "https://pypi.org"

//Provider is the plugin entry point
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"pypi_package_file":      dataSourcePackageFile(),
			"pypi_requirements_file": dataSourceRequirementsFile(),
		},
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  DefaultPyPiAddress,
				// DefaultFunc: schema.EnvDefaultFunc("PYPI_ADDR", nil),
			},
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address: d.Get("address").(string),
	}

	return config.Client()
}
