package greeterpb  // ✅ Ensure this matches the implementation package

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock GreeterServer implementation
type mockGreeterServer struct {
	UnimplementedGreeterServer
}

func (s *mockGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello " + req.Name}, nil
}

func TestSayHello(t *testing.T) {
	server := &mockGreeterServer{}

	// Test case
	req := &HelloRequest{Name: "Alice"}
	resp, err := server.SayHello(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Hello Alice", resp.Message)
}

