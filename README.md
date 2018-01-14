# kpxcconvert

[![Go Report](https://img.shields.io/badge/go%20report-A%2B-brightgreen.svg)](https://goreportcard.com/report/github.com/mdzhang/kpxcconvert)
[![GoDoc](https://godoc.org/github.com/mdzhang/kpxcconvert?status.svg)](https://godoc.org/github.com/mdzhang/kpxcconvert)
[![Build Status](https://travis-ci.org/mdzhang/kpxcconvert.svg?branch=master)](https://travis-ci.org/mdzhang/kpxcconvert)
[![GitHub release](https://img.shields.io/github/release/mdzhang/kpxcconvert.svg)](https://github.com/mdzhang/kpxcconvert)

A CLI tool for converting 1Password export files to KeePassXC import files.

Confirmed with 1Password 6.8.3 and KeePassXC 2.2.4

## Usage

Say you export your 1Password vault data to `~/Documents/1Password\ 2018-01-09,\ 08_23\ PM\ \(99\ items\ and\ 0\ folders\).1pif` for a vault called "Primary" and you want to generate a KeePassXC CSV-import compatible CSV called `out.csv`. Then you might run:

```sh
kpxcconvert --group Primary --op ~/Documents/1Password\ 2018-01-09,\ 08_23\ PM\ \(99\ items\ and\ 0\ folders\).1pif/data.1pif --kp out.csv
```

## Installation

##### For non-gophers

- Download the platform-specific tarball from the [release page](https://github.com/mdzhang/kpxcconvert/releases)
- `sudo tar xvfz kpxcconvert.tar.gz -C /usr/local/bin`
- `sudo chmod +x /usr/local/bin/kpxcconvert`

##### For gophers

```sh
go get github.com/mdzhang/kpxcconvert
```

## What it does

Will always parse 1Password titles. Username/password/URL depends on the 1Password entry type. Extra fields will always go into the `Notes` section as a yaml formatted string.

| 1Password entry type | Supported | Summary | Username       | Password            | URL             | Notes        |
| -------------------- | --------- | ------- | --------       | --------            | ---             | -----        |
| login                | O         | --      | 1pass username | 1pass password      | 1pass first URL | +extra URLs  |
| password             | O         | skipped | --             | --                  | --              | --           |
| router               | O         | --      | SSID           | 1pass wifi password | --              | --           |
| secure notes         | O         | --      | --             | --                  | --              | +1pass notes |
| credit cards         | O         | --      | cardholder     | cc number           | --              | +cc fields   |
| identities           | X         | --      | --             | --                  | --              | --           |
| licenses             | X         | --      | --             | --                  | --              | --           |
| SSN                  | X         | --      | --             | --                  | --              | --           |
| bank account         | X         | --      | --             | --                  | --              | --           |
| email account        | X         | --      | --             | --                  | --              | --           |
| driver's licenses    | X         | --      | --             | --                  | --              | --           |

## Development

### Prerequisites

- Go 1.9+
- [glide](https://glide.sh) for dependency management

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

Install package dependencies:

  ```sh
  make install
  ```

Install dev dependencies:

  ```sh
  make deps
  ```

### Tasks

Run tests

  ```sh
  make test
  ```

Lint code

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

Cut a release

- Tag a commit
- Push the tag to Travis to trigger creating a Github release
