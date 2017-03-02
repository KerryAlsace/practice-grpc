package main

import (
	"log"
	"net"
	"os"

	"github.com/KerryAlsace/practice-grpc/calculations"
	pb "github.com/KerryAlsace/practice-grpc/routes"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement practice_app.Server
type server struct{}

// AddOne implements practice_app.Server and adds 1 to the requested integer
func (s *server) AddOne(ctx context.Context, in *pb.NumberRequest) (*pb.NumberResponse, error) {
	calculations.AddOne(&in.Number)
	return &pb.NumberResponse{Number: in.Number}, nil
}

// Square implements practice_app.Server and squares the requested integer
func (s *server) Square(ctx context.Context, in *pb.NumberRequest) (*pb.NumberResponse, error) {
	calculations.Square(&in.Number)
	return &pb.NumberResponse{Number: in.Number}, nil
}

func main() {
	lis, err := net.Listen("tcp", os.Getenv("PORT"))

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
