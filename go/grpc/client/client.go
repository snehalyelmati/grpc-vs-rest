package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/snehalyelmati/grpc-vs-rest-benchmark/go/grpc/client/protos/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName     = "world"
	defaultCount    = 1
	defaultPort     = 50051
	defaultProtocol = "http"
)

var (
	name     = flag.String("name", defaultName, "Name to greet")
	count    = flag.Int("count", defaultCount, "Number of requests")
	port     = flag.Int("port", defaultPort, "Port")
	protocol = flag.String("protocol", defaultProtocol, "Protocol, options 'grpc' or 'rest'")
)

func main() {
	flag.Parse()

	diff := 0.0
	if *protocol == "grpc" {
		conn, err := grpc.Dial(":"+strconv.Itoa(*port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewHelloClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		start := float64(time.Now().UnixMilli())
		for i := 0; i < *count; i++ {
			_, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
			if err != nil {
				log.Fatalf("Could not say hello: %v", err)
			}
		}
		diff = float64(time.Now().UnixMilli()) - start
	}

	if *protocol == "http" {
		data := []byte(`
            "Name": "world"
            `)
		start := float64(time.Now().UnixMilli())
		for i := 0; i < *count; i++ {
            _, err := http.Post("http://127.0.0.1:"+strconv.Itoa(*port), "application/json; charset=utf-8", bytes.NewBuffer(data))
			if err != nil {
				log.Fatalf("Could not say hello(HTTP): %v", err)
			}
		}
		diff = float64(time.Now().UnixMilli()) - start
	}

	log.Printf("Total time taken to run %d requests: %vms", *count, float64(diff)/float64(*count))
}
