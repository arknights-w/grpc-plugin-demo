# build server
.PHONY: build
build:
	go build -o ./bootstrap/client/tencent ./bootstrap/server/

# run plugin
.PHONY: run
run:
	cd bootstrap/client/ && go run .

# build proto file to go file
.PHONY: protoc
protoc:
	protoc -I proto/ proto/text_message.proto --go_out=proto/
	protoc -I proto/ proto/text_message.proto --go-grpc_out=proto/