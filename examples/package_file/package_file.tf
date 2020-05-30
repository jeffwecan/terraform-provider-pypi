# This configuration sample shows how to download a single package archive file from a PyPi-compatabile repository
provider "pypi" {
  address = var.pypi_address
}

data "pypi_package_file" "hvac_latest" {
  name       = "hvac"
  output_dir = "${path.module}/hvac_latest"
}

data "pypi_package_file" "hvac_0-10-1" {
  name       = "hvac"
  version    = "0.10.1"
  output_dir = "${path.module}/hvac_0-10-1"
}

data "pypi_package_files" "mah_requirements" {
  requirements_file = "${path.module}/requirements.txt"
  output_dir        = "${path.module}/mah_requirements"
}
