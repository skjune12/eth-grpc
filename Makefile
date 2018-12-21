NAME     := eth-grpc
VERSION  := v0.0.1
SOLC     := solc
ABIGEN   := abigen
CONTRACT := ./contract/ExampleContract
PKGNAME  := contract

.PHONY: all
all: protoc abigen binary

.PHONY: protoc
protoc:
	protoc -I api/ --go_out=plugins=grpc:api api/api.proto

.PHONY: abigen
abigen: $(CONTRACT).sol
	abigen --sol $(CONTRACT).sol --pkg $(PKGNAME) --out=$(CONTRACT).go

.PHONY: binary
binary:
	go build -o bin/server server/main.go
	go build -o bin/client client/main.go

.PHONY: clean
clean:
	rm -f bin/server bin/client
