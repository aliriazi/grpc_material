package greeterpb

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

// MockGreeterServer implements the GreeterServer interface
type MockGreeterServer struct {
	UnimplementedGreeterServer // Embed to satisfy interface
}

// SayHello implements the gRPC method for testing
func (s *MockGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello " + req.Name}, nil
}

// StartTestServer starts a test gRPC server and returns its address
func StartTestServer() (*grpc.Server, string) {
	server := grpc.NewServer()
	mockGreeter := &MockGreeterServer{}

	// Register the gRPC service
	RegisterGreeterServer(server, mockGreeter)

	// Listen on a random port
	listener, err := net.Listen("tcp", ":0") // ":0" picks a free port
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	go server.Serve(listener) // Start the server in a goroutine

	return server, listener.Addr().String()
}

func TestGreeterServiceIntegration(t *testing.T) {
	// Start the test server
	server, addr := StartTestServer()
	defer server.Stop() // Ensure cleanup

	// Connect to the gRPC server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	// Create a client
	client := NewGreeterClient(conn)

	// Send a request
	req := &HelloRequest{Name: "Bob"}
	resp, err := client.SayHello(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Hello Bob", resp.Message)
}

