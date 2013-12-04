all: test

test:
	@go test ./...

benchmark:
	@go test -bench . ./...

coverage:
	@go get code.google.com/p/go.tools/cmd/cover
	@go test -cover ./...
