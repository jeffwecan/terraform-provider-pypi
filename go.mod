module github.com/jeffwecan/terraform-provider-pypi

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.13.0
	github.com/jeffwecan/go-pypi v0.0.1
	github.com/stretchr/testify v1.5.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/jeffwecan/go-pypi => ../go-pypi
