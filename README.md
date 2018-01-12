# kpxcconvert

A CLI tool for converting 1Password export files to KeePassXC import files.

Confirmed with 1Password 6.8.3 and KeePassXC 2.2.4

[![Go Report](https://img.shields.io/badge/go%20report-A%2B-brightgreen.svg)](https://goreportcard.com/report/github.com/mdzhang/kpxcconvert)
[![GoDoc](https://godoc.org/github.com/mdzhang/kpxcconvert?status.svg)](https://godoc.org/github.com/mdzhang/kpxcconvert)

## Usage

Say you export your 1Password vault data to `~/Documents/1Password\ 2018-01-09,\ 08_23\ PM\ \(99\ items\ and\ 0\ folders\).1pif` for a vault called "Primary" and you want to generate a KeePassXC CSV-import compatible CSV called `out.csv`. Then you might run:

```sh
kpxcconvert --group Primary --op ~/Documents/1Password\ 2018-01-09,\ 08_23\ PM\ \(99\ items\ and\ 0\ folders\).1pif/data.1pif --kp out.csv
```
## Installation

##### For non-gophers

- Download the platform-specific tarball
- `sudo tar xvfz kpxcconvert.tar.gz -C /usr/local/bin`
- `sudo chmod +x /usr/local/bin/kpxcconvert`

##### For gophers

```sh
go get github.com/mdzhang/kpxcconvert
```

## What it does

| 1Password entry type | Effects |
| -------------------- | ------- |
| login              | parse title, username, password, urls (extras go to KeePassXC notes field) |
| password             | deliberately ignored |
| router               | parse title, ssid as username, password |
| secure notes         | parse title, notes |
| credit cards         | X |
| identities         | X |
| licenses         | X |
| SSN         | X |
| bank account         | X |
| email account         | X |
| driver's licenses         | X |

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
