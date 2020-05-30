package pypi

import (
	api "github.com/jeffwecan/go-pypi/pypi"
)

// Config Used to configure an api client for PyPi
type Config struct {
	Address  string
}

// Client Returns a new PyPi api client configured with this instance parameters
func (c *Config) Client() (*api.PackageIndex, error) {
	return api.NewPackageIndex(c.Address), nil
}
