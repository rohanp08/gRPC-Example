package main

import (
	"context"
	pb "gRPC/student"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "Student1"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//request ID
	r, err := c.StudentID(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("could not get ID: %v", err)
	}
	log.Printf("Student: %v, ID: %v",name, r.ID)
}


