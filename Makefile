export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

ONOS_PROTOC_VERSION := v0.6.3

test: # @HELP run the unit tests and source code validation
test: deps license_check linters
	go test github.com/onosproject/onos-ric-sdk-go/pkg/...

coverage: # @HELP generate unit test coverage data
coverage: deps linters license_check
	# ./build/bin/coveralls-coverage

deps: # @HELP ensure that the required dependencies are in place
	GOPRIVATE="github.com/onosproject/*" go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION}

all: test

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
