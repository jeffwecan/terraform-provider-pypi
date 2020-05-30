package pypi_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourcePackageFile_Name(t *testing.T) {
	resName := "data.pypi_package_file.name"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProjectName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", "some-python-package"),
				),
			},
		},
	})
}


const testAccDataSourceProjectName = `
data "pypi_package_file" "name" {
  name = "some-python-package"
}
`
