build:
	@go build -o bin/ecom-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-go