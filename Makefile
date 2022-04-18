install:
	go mod download
	go mod verify
	go mod tidy

protos:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/

clean:
	rm -rf ./gen/proto/*.go