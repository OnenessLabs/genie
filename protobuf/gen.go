//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. api/api.proto
package protobuf
