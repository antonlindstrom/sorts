all: test

test:
	@go test ./...

benchmark:
	@go test -bench . ./...
