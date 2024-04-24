init: install add-proto

install:
	go mod tidy

add-proto:
	protoc --go_out=. --go-grpc_out=. proto/calculator.proto

run-server:
	go run server/main.go

run-client:
	go run client/main.go