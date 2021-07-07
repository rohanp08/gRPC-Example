package main

import (
	"context"
	pb "gRPC/student"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement student.StudentID
type server struct {
	pb.UnimplementedGreeterServer
}

// StudentID implements student.StudentID
func (s *server) StudentID(ctx context.Context, in *pb.Request) (*pb.ReplyID, error) {
	log.Printf("Received: %v", in)
	id := rand.Int63() //generate random ID
	reply := &pb.ReplyID{ID: id}
	return reply, nil //return reply which contains ID
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
