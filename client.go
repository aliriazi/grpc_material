package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"github.com/aliriazi/grpc_material/greeterpb" // import the generated package
)

func main() {
	// Set up the connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new Greeter client
	client := greeterpb.NewGreeterClient(conn)

	// Create a request object
	req := &greeterpb.HelloRequest{Name: "World"}

	// Call the SayHello RPC method
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}

	// Print the response
	fmt.Printf("Response: %s\n", res.GetMessage())
}

