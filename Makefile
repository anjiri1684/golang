build:
	@go build - bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom
