package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	."protocol"
	"errors"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Add(ctx context.Context, in *CalParam) (*CalResult, error) {
	result := in.A + in.B
	return &CalResult{Result:result}, nil
}
func (s *server) Sub(ctx context.Context, in *CalParam) (*CalResult, error) {
	result := in.A - in.B
	return &CalResult{Result:result}, nil
}
func (s *server) Mul(ctx context.Context, in *CalParam) (*CalResult, error) {
	result := in.A * in.B
	return &CalResult{Result:result}, nil
}
func (s *server) Div(ctx context.Context, in *CalParam) (*CalResult, error) {
	if 0 == in.B {
		err := errors.New("The divisor cannot be zero!")
		return nil, err
	}
	result := in.A / in.B
	return &CalResult{Result:result}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterCalculatorServer(s, &server{})
	// Register reflection service on gRPC server.
	//reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
