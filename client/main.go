package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/skjune12/grpc-eth/api"
	"google.golang.org/grpc"
)

func main() {
	// setup gRPC conn
	conn, err := grpc.Dial("localhost:4567", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial: %s\n", err)
	}
	defer conn.Close()

	// setup gRPC client
	grpcClient := api.NewExampleClient(conn)

	if os.Args[1] == "add" {
		value, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		response, err := grpcClient.Exec(context.Background(), &api.TestMsg{Method: api.ADD, Value: int32(value)})
		if err != nil {
			log.Fatalf("Error when calling grpcClient.Exec: %s\n", err)
		}

		log.Printf("Response from server: %s\n", response.Msg)
	}

	if os.Args[1] == "get" {
		response, err := grpcClient.Exec(context.Background(), &api.TestMsg{Method: api.GET})
		if err != nil {
			log.Fatalf("Error when calling grpcClient.Exec: %s\n", err)
		}

		log.Printf("Response from server: %s\n", response.Msg)
	}
}
