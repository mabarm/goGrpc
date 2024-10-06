package main

import (
	"context"

	pb "github.com/mabarm/goGrpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello Mamata",
	}, nil
}
