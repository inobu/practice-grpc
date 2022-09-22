.PHONY: proto

proto: proto
	protoc -I=./proto --go_out=./ --go-grpc_out=./ ./proto/*.proto --go-grpc_opt require_unimplemented_servers=false
