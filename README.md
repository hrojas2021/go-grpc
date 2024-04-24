## GRPC with complex data types

This is a template for basic GRPC communication client -> server -> client.\
This example will simulate a calculator with **add** and **substract** operation.
There's an extra Complex function that will have a bunch of different supported data types in proto.

## Initial Setup

Install the basic protoc of golang to build the proto files

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Run basic commands to install dependencies and build proto files

```
go mod tidy
add-proto
```

## Server

```
make run-server
```

## Client

```
make run-server
```
