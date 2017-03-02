package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement practice_app.Server
type server struct{}

// AddOne implements practice_app.Server and adds 1 to the requested integer
func (s *server) AddOne(ctx context.Context, in *pb.NumberRequest) (*pb.NumberReply, error) {
	calculations.AddOne(&in.Number)
	return &pb.NumberResponse{Message: in.Number}, nil
}

// Square implements practice_app.Server and squares the requested integer
func (s *server) Square(ctx context.Context, in *pb.NumberRequest) (*pb.NumberReply, error) {
	calculations.Square(&in.Number)
	return &pb.NumberResponse{Message: in.Number}, nil
}

func main() {
	lis, err := net.Listen("tcp", os.Getenv(PORT))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
