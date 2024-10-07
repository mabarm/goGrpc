package main

import (
	"log"

	pb "github.com/mabarm/goGrpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	callSayHello(client)

}

/*
Explanation of Key Parts:
1. log package: Used to log errors and other information.
2. gRPC libraries: Used to handle gRPC communication and transport security.
3. port = ":8080": Specifies the port where the gRPC server is running.
4. grpc.NewClient: Establishes a connection to the server. Here, it's using an insecure transport (non-TLS) because of insecure.NewCredentials().
5. client := pb.NewGreetServiceClient(conn): Creates a new client for the GreetService, a service generated from the protobuf file.
6. callSayHello(client): A placeholder function that will handle calling the SayHello RPC method using the created client.
*/
