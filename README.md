# IT RALLY 2020 HighLoad Cup: The task

[![CI/CD](https://github.com/Djarvur/allcups-itrally-2020-task/workflows/CI/CD/badge.svg?event=push)](https://github.com/Djarvur/allcups-itrally-2020-task/actions?query=workflow%3ACI%2FCD)
[![CircleCI](https://circleci.com/gh/Djarvur/allcups-itrally-2020-task.svg?style=svg&circle-token=245b43b1cdcb425be9eaa937cc2ae54b88d54dc9)](https://circleci.com/gh/Djarvur/allcups-itrally-2020-task)
[![Project Layout](https://img.shields.io/badge/Standard%20Go-Project%20Layout-informational)](https://github.com/golang-standards/project-layout)

Service implementing the task for IT RALLY 2020 HighLoad Cup (runs on All Cups platform).

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Development](#development)
  - [Requirements](#requirements)
  - [Setup](#setup)
  - [Usage](#usage)
- [Deploy](#deploy)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Development

### Requirements

- Go 1.15
- [Docker](https://docs.docker.com/install/) 19.03+
- [Docker Compose](https://docs.docker.com/compose/install/) 1.25+
- Tools used to build/test project (feel free to install these tools using
  your OS package manager or any other way, but please ensure they've
  required versions; also note these commands will install some non-Go
  tools into `$GOPATH/bin` for the sake of simplicity):

```sh
curl -sSfL https://github.com/hadolint/hadolint/releases/download/v1.18.0/hadolint-$(uname)-x86_64 | install /dev/stdin $(go env GOPATH)/bin/hadolint
curl -sSfL https://github.com/koalaman/shellcheck/releases/download/v0.7.1/shellcheck-v0.7.1.$(uname).x86_64.tar.xz | tar xJf - -C $(go env GOPATH)/bin --strip-components=1 shellcheck-v0.7.1/shellcheck
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
go get gotest.tools/gotestsum@v0.5.3
go get github.com/golang/mock/mockgen@v1.4.4
go get github.com/cheekybits/genny@master
curl -sSfL https://github.com/go-swagger/go-swagger/releases/download/v0.25.0/swagger_$(uname)_amd64 | install /dev/stdin $(go env GOPATH)/bin/swagger
```

### Setup

1. After cloning the repo copy `env.sh.dist` to `env.sh`.
2. Review `env.sh` and update for your system as needed.
3. It's recommended to add shell alias `alias dc="if test -f env.sh; then
   source env.sh; fi && docker-compose"` and then run `dc` instead of
   `docker-compose` - this way you won't have to run `source env.sh` after
   changing it.

### Usage

To develop this project you'll need only standard tools: `go generate`,
`go test`, `go build`, `docker build`. Provided scripts are for
convenience only.

- Always load `env.sh` *in every terminal* used to run any project-related
  commands (including `go test`): `source env.sh`.
    - When `env.sh.dist` change (e.g. by `git pull`) next run of `source
      env.sh` will fail and remind you to manually update `env.sh` to
      match current `env.sh.dist`.
- `go generate ./...` - do not forget to run after making changes related
  to auto-generated code
- `go test ./...` - test project (excluding integration tests), fast
- `./scripts/test` - thoroughly test project, slow
- `./scripts/test-ci-circle` - run tests locally like CircleCI will do
- `./scripts/cover` - analyse and show coverage
- `./scripts/build` - build docker image and binaries in `bin/`
    - Then use mentioned above `dc` (or `docker-compose`) to run and
      control the project.
    - Access project at host/port(s) defined in `env.sh`.

As this project isn't a real service but a _one-shot task_ which is
supposed to handle single user for a fixed period of time and then finish,
if you'll use `dc up` to start it while development then you should use
`dc up --force-recreate` to ensure each time it'll start with clean state.
An alternative is to just run `bin/task` or use `docker run`.

## Deploy

```
docker run --name=hlcup2020_task -i -t --rm \
    -e HLCUP2020_DIFFICULTY=normal \
    -v hlcup2020-task:/home/app/var/data \
    ghcr.io/djarvur/allcups-itrally-2020-task:1.0.0
```
