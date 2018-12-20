NAME     := eth-grpc
VERSION  := v0.0.1
SOLC     := solc
ABIGEN   := abigen
CONTRACT := ./contract/ExampleContract
PKGNAME  := contract

.PHONY: all
all: protoc abigen

.PHONY: protoc
protoc:
	protoc -I api/ --go_out=plugins=grpc:api api/api.proto

.PHONY: abigen
abigen: $(CONTRACT).sol
	abigen --sol $(CONTRACT).sol --pkg $(PKGNAME) --out=$(CONTRACT).go
