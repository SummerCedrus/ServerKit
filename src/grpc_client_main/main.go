
package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer func() {
		conn.Close()
		cancel()
	}()
	c := NewCalculatorClient(conn)

	// Contact the server and print out its response.

	r, err := c.Add(ctx, &CalParam{A:3,B:2})
	if err != nil {
		log.Fatalf("could not cal: %v", err)
	}
	log.Printf("Result: %d", r.GetResult())
}
