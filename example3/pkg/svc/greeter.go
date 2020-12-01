package svc

import (
	"context"
	"log"

	"github.com/ceperapl/go-bazel-tutorial/example3/pkg/pb"
)

type GreeterServer struct {
	pb.GreeterServer
}

// SayHello implements pb.GreeterServer
func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
