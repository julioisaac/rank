SHELL=/bin/bash

run:
	go run main.go $(path)

tests:
	go test -v -cover `go list ./...`