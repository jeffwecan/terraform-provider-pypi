# terraform-provider-pypi

Terraform provider aimed at facilitating downloading Python requirements when needed during Terraform runs. For example, when packaging an AWS lambda deployment.

## Usage

Downloading the latest release of a package to a directory (this will attempt to download a `bdist_wheel` artifact by default and fallback to a `sdist` only if the `*whl` file isn't available):

```hcl
data "pypi_package_file" "hvac_latest" {
  name       = "hvac"
  output_dir = "${path.module}/hvac_latest"
}
```

Downloading a specific version:

```hcl
data "pypi_package_files" "mah_requirements" {
  requirements_file = "${path.module}/requirements.txt"
  output_dir = "${path.module}/hvac_0-10-1"
}
```

## Development

### Installing Locally

```shellsession
$ VERSION=0.0.1 make install
GO111MODULE=on go build -o ./bin/terraform-provider-pypi_0.0.1
cp ./bin/terraform-provider-pypi_0.0.1 ~/.terraform.d/plugins/
```

### Releases

```shellsession
# Hit CTRL+d after typing some stuff to cache your GPG key password...
$ gpg --armor --detach-sign
hi
-----BEGIN PGP SIGNATURE-----
[...]
-----END PGP SIGNATURE-----
```

```shellsession
$ export GITHUB_TOKEN='<personal access token with public_repo scope>'
$ git tag v0.0.1
$ goreleaser release --rm-dist
```
