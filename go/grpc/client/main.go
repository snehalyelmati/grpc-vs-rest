package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/snehalyelmati/grpc-vs-rest-benchmark/go/grpc/client/protos/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName  = "world"
	defaultCount = 1
)

var (
	port  = ":50051"
	name  = flag.String("name", defaultName, "Name to greet")
	count = flag.Int("count", defaultCount, "Number of requests")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 0; i < *count; i++ {
		start := time.Now().UnixMilli()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("Could not say hello: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		diff := time.Now().UnixMilli() - start
		log.Printf("Time taken to run the request: %vms", diff)
	}
	// log.Printf("Time taken to run %d requests: %vms", count, diff)
}
