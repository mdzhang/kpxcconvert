# kpxcconvert

A CLI tool for converting 1Password export files to KeePassXC import files

## Installation

- Download the platform-specific tarball
- `sudo tar xvfz kpxcconvert.tar.gz -C /usr/local/bin`
- `sudo chmod +x /usr/local/bin/kpxcconvert`

## Development

### Prerequisites

- Go
- golint - `go get -u github.com/golang/lint/golint`
- gox - `go get -u github.com/mitchellh/gox`
- glide - go dependency manager
- delve - `go get -u github.com/derekparker/delve/cmd/dlv`

### Setup

Create and enter package directory

  ```sh
  mkdir -p $GOPATH/src/gitub.com/mdzhang
  cd $GOPATH/src/gitub.com/mdzhang
  ```

Clone kpxcconvert

  ```sh
  git clone git@github.com:mdzhang/kpxcconvert.git
  ```

Install dependencies:

  ```sh
  glide install
  ```

### Tasks

Run tests

  ```sh
  make test
  ```

Lint code:

  ```sh
  make lint
  ```

Compile and generate binary for current platform/architecture

  ```sh
  make compile
  ```

Cross-platform packaging

  ```sh
  make package
  ```
