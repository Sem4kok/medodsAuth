tidy:
	@go mod tidy
build:
	@go build -o ./bin/app ./cmd/main.go
run: tidy build
	@./bin/app
test:
	@go test -v ./...
start:
	@sudo docker-compose up --build
clear:
	@sudo docker-compose down