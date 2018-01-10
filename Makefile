OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
BUILD_REF ?= $(shell git rev-parse --verify HEAD)
VERSION := $(shell git describe --tags --abbrev=0)
PROJECT_NAME=kpxcconvert

build: compile

compile:
	@echo "$(OK_COLOR)==> Compiling binary$(NO_COLOR)"
	@go build -o bin/${PROJECT_NAME}

compile-all:
	@gox -verbose \
	-ldflags "-X main.version=${VERSION}" \
	-os="linux darwin" \
	-arch="amd64" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

package: compile-all
	@tar -cvzf dist/${PROJECT_NAME}-${VERSION}-darwin-amd64.tar.gz -C dist/darwin-amd64 .
	@tar -cvzf dist/${PROJECT_NAME}-${VERSION}-linux-amd64.tar.gz -C dist/linux-amd64 .

history:
	@git changelog --tag $(VERSION)

clean:
	@rm -f ./bin
	@rm -rf ./dist

lint:
	@golint *.go
	@golint cli/
	@golint logger/
	@golint version/

format:
	@gofmt -w .

sloc:
	@wc -l *.go */*.go

test: lint
	@go test -v -race cli/*.go
	@go test -v -race logger/*.go
	@go test -v -race version/*.go

install:
	@glide install
