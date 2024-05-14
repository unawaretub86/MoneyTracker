# run test coverage
test:
	go test -v -cover ./...
# run linter run it to avoid waste actions minuts
lint:
	golangci-lint run
# start the application
start:
	go run cmd/main.go
#rm if u need use docker	
#docker-compose up