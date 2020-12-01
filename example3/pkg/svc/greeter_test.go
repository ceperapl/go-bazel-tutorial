package svc

import (
	"context"
	"reflect"
	"testing"

	"github.com/ceperapl/go-bazel-tutorial/example3/pkg/pb"
)

func TestSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  *pb.HelloRequest
		response *pb.HelloReply
		err      error
	}{
		{
			"test 1",
			&pb.HelloRequest{Name: "cepera"},
			&pb.HelloReply{Message: "Hello cepera"},
			nil,
		}, {
			"test 2",
			&pb.HelloRequest{Name: "guest"},
			&pb.HelloReply{Message: "Hello guest"},
			nil,
		}, {
			"test 3",
			&pb.HelloRequest{Name: ""},
			&pb.HelloReply{Message: "Hello "},
			nil,
		},
	}

	server := GreeterServer{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := server.SayHello(context.Background(), tt.request)
			if err != tt.err {
				t.Errorf("\nactual: %q\nexpected: %q\n", err, tt.err)
			}

			if !reflect.DeepEqual(response, tt.response) {
				t.Errorf("\nactual: %v\nexpected: %v\n", response, tt.response)
			}
		})
	}
}
