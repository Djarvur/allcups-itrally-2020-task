# IT RALLY 2020 HighLoad Cup: The task

[![CI/CD](https://github.com/Djarvur/allcups-itrally-2020-task/workflows/CI/CD/badge.svg?event=push)](https://github.com/Djarvur/allcups-itrally-2020-task/actions?query=workflow%3ACI%2FCD) [![CircleCI](https://circleci.com/gh/Djarvur/allcups-itrally-2020-task.svg?style=svg&circle-token=245b43b1cdcb425be9eaa937cc2ae54b88d54dc9)](https://circleci.com/gh/Djarvur/allcups-itrally-2020-task) ![Go](https://github.com/Djarvur/allcups-itrally-2020-task/workflows/Go/badge.svg?event=push) [![Project Layout](https://img.shields.io/badge/Standard%20Go-Project%20Layout-informational)](https://github.com/golang-standards/project-layout) [![Release](https://img.shields.io/github/v/release/Djarvur/allcups-itrally-2020-task)](https://github.com/Djarvur/allcups-itrally-2020-task/releases/latest)

Service implementing the task for IT RALLY 2020 HighLoad Cup (runs on All Cups platform).

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
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

#### Cheatsheet

```sh
dc up -d --remove-orphans               # (re)start all project's services
dc logs -f -t                           # view logs of all services
dc logs -f SERVICENAME                  # view logs of some service
dc ps                                   # status of all services
dc restart SERVICENAME
dc exec SERVICENAME COMMAND             # run command in given container
dc stop && dc rm -f                     # stop the project
docker volume rm PROJECT_SERVICENAME    # remove some service's data
```

It's recommended to avoid `docker-compose down` - this command will also
remove docker's network for the project, and next `dc up -d` will create a
new network… repeat this many enough times and docker will exhaust
available networks, then you'll have to restart docker service or reboot.

## Deploy

TODO

```
docker run --rm ghcr.io/Djarvur/allcups-itrally-2020-task
```
