OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
BUILD_REF ?= $(shell git rev-parse --verify HEAD)
VERSION := $(shell git describe --tags --abbrev=0)
PROJECT_NAME=kpxcconvert
SRC = ". ./cli/ ./logger/ ./version/ ./kpxc/ ./opass/ ./secret/"

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
	@echo ${SRC} | xargs -n1 golint

vet:
	@echo ${SRC} | xargs -n1 go vet -x

format:
	@gofmt -w .

sloc:
	@wc -l *.go */*.go

test:
	@echo ${SRC} | xargs -n1 go test -race

install:
	@glide install

deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/mitchellh/gox
	go get -u github.com/derekparker/delve/cmd/dlv
	go get -u github.com/c4milo/github-release

release: package
	./release.sh
