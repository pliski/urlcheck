# urlcheck

CLI tool for checking the http responses for a given list of URLs.

## Install

One-line remote install (`CGO_ENABLED=0` builds a static binary so it runs on
NixOS too):

```sh
CGO_ENABLED=0 go install github.com/pliski/urlcheck@latest
```

Or build from a clone:

```sh
git clone git@github.com:pliski/urlcheck.git
cd urlcheck
make install
```

Either way the binary lands in `$GOBIN` (defaults to `$(go env GOPATH)/bin`);
make sure that directory is on your `PATH`.

The `CGO_ENABLED=0` matters on NixOS: a plain `go install` produces a
dynamically linked binary that fails with "Could not start dynamically linked
executable". See [Build](#build) for details.

Usage:

```sh
$ urlcheck https://go.dev/ https://charm.sh/
$ urlcheck -f ~/urls.txt
```

## Inline and script usage

Check if go.dev is responding and wait max 20 seconds for the response:

```sh
#!/bin/bash 

if urlcheck -s http://go.dev -t 20 1>/dev/null 2>&1; then
  echo go.dev is UP
else
  echo go.dev is DOWN
fi
```

## Build

Requires [Go](https://go.dev/) and `make`.

```sh
make build                              # static binary -> ./bin/urlcheck
make install                            # static binary -> $GOBIN (go install)
make dynamic                            # cgo/dynamic binary -> ./bin/urlcheck
make run ARGS="-s http://go.dev -t 20"  # build (static) and run
make clean                              # remove ./bin
```

The default build sets `CGO_ENABLED=0` to produce a statically linked,
dependency-free binary. That is what lets it run on NixOS, which does not ship
the `/lib64/ld-linux-x86-64.so.2` interpreter a normal cgo build hardcodes. The
static binary is also portable to musl/Alpine and scratch containers.

Use `make dynamic` only if you need the cgo DNS resolver (it honors
`nsswitch.conf` / NSS modules); the resulting binary is tied to the host it was
built on.

## Dependencies

* [bubble tea](https://github.com/charmbracelet/bubbletea/tree/master/tutorials/commands/)

* [cobra](https://cobra.dev/)

## Todos 

- [ ] use timeoutRoundTripper instead of http.Cliet.Get also for urlist checking in urlcheck.go
- [ ] -s flag : check also for the whole urilist and exit at the first fail
- [ ] feat: continuous mode with retry interval


