#!/usr/bin/env bash

# fail automatically as soon as an error is detected
set -e

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16.2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.16.2
