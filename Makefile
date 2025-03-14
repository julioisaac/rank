SHELL=/bin/bash
GOPACKAGES=$(shell go list ./... | egrep -v "data|main.go")

run:
	go run main.go $(path)

tests:
	go test -v -cover `go list ./...`

clean-coverage:
	rm -rf cover.out

coverage: clean-coverage
	ENVIRONMENT=test go test ./... -covermode="count" -coverprofile="cover.out" $(GOPACKAGES)
