# Build settings for urlcheck.
#
# CGO_ENABLED=0 produces a statically linked, pure-Go binary: no libc
# dependency and no hardcoded ELF interpreter (/lib64/ld-linux-x86-64.so.2).
# This is required to run on NixOS, which does not ship that interpreter, and
# also makes the binary portable to musl/Alpine and scratch containers.
#
# Exporting it here means every target below builds statically by default.
export CGO_ENABLED := 0

BINARY  := urlcheck
BIN_DIR := bin

# -s strips the symbol table, -w strips DWARF debug info (smaller binary).
LDFLAGS := -s -w

.PHONY: all build dynamic install run clean

all: build

## build: compile a static binary into ./bin (default; runs anywhere, incl. NixOS)
build:
	go build -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY) .

## dynamic: compile a dynamically linked (cgo) binary into ./bin.
##          Enables the cgo DNS resolver (honors nsswitch.conf / NSS modules),
##          but the binary is tied to the glibc loader path baked in at build
##          time: built on NixOS it points into /nix/store (runs here, not
##          portable); built on a generic distro it points at /lib64 (won't run
##          on NixOS without nix-ld). Use `build` (static) for a portable binary.
dynamic: export CGO_ENABLED := 1
dynamic:
	go build -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY) .

## install: install a static binary into GOBIN (go install)
install:
	go install -ldflags="$(LDFLAGS)" .

## run: build then run, e.g. make run ARGS="-s http://go.dev -t 20"
run: build
	./$(BIN_DIR)/$(BINARY) $(ARGS)

## clean: remove build artifacts
clean:
	rm -rf $(BIN_DIR)
