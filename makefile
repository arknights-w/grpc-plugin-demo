# build server
.PHONY: build
build:
	go build -o ./bootStrap/client/tencent ./bootStrap/server/

# run plugin
.PHONY: run
run:
	cd bootStrap/client/ && go run .