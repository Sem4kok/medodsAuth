build:
	@go build -o bin/medodsAuth
run: build
	@./bin/medodsAuth
test:
	@go test -v ./...

