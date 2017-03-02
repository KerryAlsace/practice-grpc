package main

import (
	"log"
	"os"
	"strconv"

	pb "github.com/KerryAlsace/practice-grpc/routes"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const defaultNum = int32(2)

func main() {
	// Set up connection to the server
	conn, err := grpc.Dial(os.Getenv("ADDRESS"), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	// Contact the server and print out its response
	num := defaultNum
	if len(os.Args) > 1 {
		n, _ := strconv.ParseInt(os.Args[1], 0, 32)
		num = int32(n)
	}

	// AddOne
	r, err := c.AddOne(context.Background(), &pb.NumberRequest{Number: num})
	if err != nil {
		log.Fatalf("Could not AddOne: %v", err)
	}
	log.Printf("AddOne: %d", r.Number)

	// Square
	r, err = c.Square(context.Background(), &pb.NumberRequest{Number: num})
	if err != nil {
		log.Fatalf("Could not Square: %v", err)
	}
	log.Printf("Square: %d", r.Number)
}
