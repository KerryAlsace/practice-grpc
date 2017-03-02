package main

import (
	"context"
	"log"
	"os"

	pb "github.com/KerryAlsace/practice-grpc/routes"

	"google.golang.org/grpc"
)

const defaultNum = 2

func main() {
	// Set up connection to the server
	conn, err := grpc.Dial(os.Getenv(ADDRESS), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	// Contact the server and print out its response
	num := defaultNum
	if len(os.Args) > 1 {
		num = os.Args[1]
	}

	// AddOne
	r, err := c.AddOne(context.Background(), &pb.NumberRequest{Number: num})
	if err != nil {
		log.Fatalf("Could not AddOne: %v", err)
	}
	log.Printf("AddOne: %d", r.Message)

	// Square
	r, err = c.Square(context.Background(), &pb.NumberRequest{Number: num})
	if err != nil {
		log.Fatalf("Could not Square: %v", err)
	}
	log.Printf("Square: %d", r.Message)
}
