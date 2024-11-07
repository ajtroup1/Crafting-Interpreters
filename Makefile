build:
	@go build -o bin/clear main.go

repl: build
	@./bin/clear

test:
	@go test -v ./...

fmt:
	@go fmt ./...