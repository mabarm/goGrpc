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

/*
Key points in the documentation:
Function Purpose:

The SayHello method is a gRPC service implementation that sends back a static greeting message.
Parameters:

ctx: The context.Context parameter manages deadlines, cancellation signals, and metadata across RPC calls.
req: This is the input message, in this case, NoParam, indicating that the RPC doesn't require any specific parameters.
Return Values:

The method returns a HelloResponse message with a static greeting and an error. Here, the error is nil since this is a simple method with no failure conditions.
*/
