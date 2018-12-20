package main

import (
	"log"
	"net"

	"github.com/skjune12/grpc-eth/api"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:4567"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listen on port :4567")

	s := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterExampleServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err)
	}
}
