build:
	@go build -o bin/goFileShare

run: build
	@./bin/goFileShare

test:
	@go test -v ./...