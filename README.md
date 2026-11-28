# stackcheck

A concurrent HTTP health checker CLI for multi-service stacks.

## Features

* Checks run concurrently with a worker pool
* YAML config
* Exits 1 if anything is down (works in scripts and CI)

## Install

```sh
go install github.com/colby-noah/stackcheck@latest
```

Make sure `$(go env GOPATH)/bin` is on your `PATH`.

or

```sh
git clone git@github.com:colby-noah/stackcheck.git
cd stackcheck
go build
```

## Usage

Create a `targets.yaml` in the directory you run it from:

```yaml
targets:
  - name: Asteroids
    url: https://colby-noah.github.io/asteroids/
  - name: Broken
    url: https://does-not-exist.invalid/
  - name: Google
    url: https://www.google.com/
```

Run it:

```sh
stackcheck
```

(or `./stackcheck` if you built from source)

Output:

```text
Asteroids       UP     200 325ms
Broken          DOWN   Get "https://does-not-exist.invalid/": dial tcp: lookup does-not-exist.invalid: no such host 2ms
Google          UP     200 205ms
```

Exits with status 1 if any target is down.

## Development

```sh
go test -race ./...
golangci-lint run
```
