package main

import (
	"context"
	"log"
	"net"

	pb "github.com/snehalyelmati/grpc-vs-rest-benchmark/go/grpc/server/protos/hello"
	"google.golang.org/grpc"
)

var (
	port = ":50051"
)

// server used to implement hello server
type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
