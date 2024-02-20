build:
	@go build -o bin/goFileShare ./cmd/goFileShare


run: build
	@./bin/goFileShare

test:
	@go test -v ./...