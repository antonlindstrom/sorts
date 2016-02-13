all: test

test:
	@go test ./...

benchmark:
	@go test -bench . ./...

coverage:
	@go get golang.org/x/tools/cmd/cover
	@go test -cover ./...
