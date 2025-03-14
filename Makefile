SHELL=/bin/bash

tests:
	go test -v -cover `go list ./...`