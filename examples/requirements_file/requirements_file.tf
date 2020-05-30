# This configuration sample shows how to download a single package archive file from a PyPi-compatabile repository
provider "pypi" {
  address = var.pypi_address
}

data "pypi_requirements_file" "requirements" {
  requirements_file = "${path.module}/requirements.txt"
  output_dir        = "${path.module}/mah_requirements"
}

data "archive_file" "peer_removal" {
  type        = "zip"
  source_dir =  "${path.module}/mah_requirements"
  output_path = "${path.module}/python_lambda.zip"
}
