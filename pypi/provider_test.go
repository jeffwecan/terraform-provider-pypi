package pypi_test

import (
	"os"
	"testing"

	"github.com/jeffwecan/terraform-provider-pypi/pypi"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = pypi.Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"pypi": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := pypi.Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("Error: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = pypi.Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("PYPI_ADDR"); v == "" {
		t.Fatal("PYPI_ADDR must be set for acceptance tests")
	}
}
