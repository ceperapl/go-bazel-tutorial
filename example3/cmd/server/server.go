package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ceperapl/go-bazel-tutorial/example3/pkg/pb"
	"github.com/ceperapl/go-bazel-tutorial/example3/pkg/svc"
	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	defaultServerPort  = "9090"
	defaultGatewayPort = "8080"
)

func main() {
	doneC := make(chan error)

	serverPort := flag.String("server.port", defaultServerPort, "port of gRPC server")
	gatewayPort := flag.String("gateway.port", defaultGatewayPort, "port of gateway http server")

	flag.Parse()

	lis, err := net.Listen("tcp", ":"+*serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &svc.GreeterServer{})

	log.Printf("Running gRPC Server (port: %s)...\n", *serverPort)
	go func() {
		doneC <- s.Serve(lis)
	}()

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", *serverPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't contact grpc server: %v", err)
	}

	mux := grpc_runtime.NewServeMux()
	err = pb.RegisterGreeterHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("Cannot serve http api: %v", err)
	}

	log.Printf("Running HTTP Server (port: %s)...\n", *gatewayPort)
	go func() {
		doneC <- http.ListenAndServe(":"+*gatewayPort, mux)
	}()

	if err := <-doneC; err != nil {
		log.Fatal(err)
	}
}
