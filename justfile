
# Commands for hellorest
default:
  @just --list
# Build hellorest with Go
build:
  go build ./...

# Run tests for hellorest with Go
test:
  go clean -testcache
  go test ./...