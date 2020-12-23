package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ceperapl/go-bazel-tutorial/example3/pkg/pb"
	"google.golang.org/grpc"
)

const (
	defaultServerAddress = "localhost"
	defaultServerPort    = "9090"
	defaultName          = "world"
)

func main() {
	serverAddress := flag.String("server.address", defaultServerAddress, "address of gRPC server")
	serverPort := flag.String("server.port", defaultServerPort, "port of gRPC server")
	name := flag.String("name", defaultName, "name to print out")
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", *serverAddress, *serverPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
