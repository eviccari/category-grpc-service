create_proto:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
	go mod download
	go mod verify
	go mod tidy

clean:
	rm -rf ./gen/proto/*.go