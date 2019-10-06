# CO Air Quality API
This is the JSON/REST API for the CO air quality tool, written in Go.

## Build
In this directory, eval `go build ./src/main.go`

## Run
In this directory, eval `./main`

## Test
I like to keep tests in a subdirectory named `test` within the same tree as the source code.
This keeps the production code and testing code close together, but doesn't clutter the
production packages.
To run these tests, eval `go test ./.../test`
