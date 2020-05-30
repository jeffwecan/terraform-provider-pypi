export GO111MODULE=on

default: build

build:
	GO111MODULE=on go build -o ./bin/registry.terraform.io/jeffwecan/pypi/${VERSION}/darwin_amd64/terraform-provider-pypi

install: build
	cp ./bin/terraform-provider-pypi_${VERSION} ~/.terraform.d/plugins/

clean:
	rm -rf ./bin

clean_examples:
	find ./examples -name '*.tfstate' -delete
	find ./examples -name ".terraform" -type d -exec rm -rf "{}" \;
