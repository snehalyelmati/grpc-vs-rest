package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	pb "github.com/snehalyelmati/grpc-vs-rest-benchmark/go/grpc/client/protos/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName  = "world"
	defaultCount = 1
	defaultPort  = 50051
)

var (
	name  = flag.String("name", defaultName, "Name to greet")
	count = flag.Int("count", defaultCount, "Number of requests")
	port  = flag.Int("port", defaultPort, "Port")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(":"+strconv.Itoa(*port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	start := time.Now().UnixMilli()
	for i := 0; i < *count; i++ {
		_, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("Could not say hello: %v", err)
		}
	}
	diff := time.Now().UnixMilli() - start
	log.Printf("Total time taken to run %d requests: %vms", *count, diff)
}
