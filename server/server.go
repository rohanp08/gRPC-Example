package main

import (
	"context"
	pb "gRPC/student"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var student_id_map = make(map[string]int64)

// server is used to implement student.StudentID
type server struct {
	pb.UnimplementedGreeterServer
}

// StudentID implements student.StudentID
func (s *server) StudentID(ctx context.Context, in *pb.Request) (*pb.ReplyID, error) {
	log.Printf("Received: %v", in)
	id := rand.Int63()
	name := in.GetName()
	if student_id_map[name] == 0 {
		id = rand.Int63()
		student_id_map[name] = id
	} else {
		id = int64(student_id_map[name])
	}
	//generate random ID
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
