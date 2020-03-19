package main

import (
	"context"
	"crawler/rpc"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewCrawlerWorkClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Process(ctx, &rpc.WorkerRequest{
		Url:  "http://www.zhenai.com/zhenghun",
		Type: "Region",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetRequests())
}
