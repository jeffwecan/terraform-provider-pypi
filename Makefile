export GO111MODULE=on

default: build

build:
	GO111MODULE=on go build -o ./bin/terraform-provider-pypi_${VERSION}

install: build
	cp ./bin/terraform-provider-pypi_${VERSION} ~/.terraform.d/plugins/

clean:
	rm -rf ./bin

clean_examples:
	find ./examples -name '*.tfstate' -delete
	find ./examples -name ".terraform" -type d -exec rm -rf "{}" \;
