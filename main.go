package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "github.com/aliriazi/grpc_material/greeterpb"
)

type server struct {
    greeterpb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *greeterpb.HelloRequest) (*greeterpb.HelloResponse, error) {
    return &greeterpb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    greeterpb.RegisterGreeterServer(grpcServer, &server{})

    log.Println("gRPC server listening on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

