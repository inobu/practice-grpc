.PHONY: proto

proto: proto
	protoc -I=./proto --go_out=./ --go-grpc_out=./ ./proto/*.proto
