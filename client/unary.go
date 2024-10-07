package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mabarm/goGrpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("Not able to greet %v", err)
	}
	log.Printf("%s", res.Message)
}

/*
Explanation of Key Parts:
1. context.WithTimeout: Creates a context with a 1-second timeout. This context is passed to the gRPC request, ensuring it doesn't wait forever if the server doesn't respond.
2. defer cancel(): This ensures that the context's cancellation function is called once the operation is complete, freeing resources even if the timeout isn't reached.
3. client.SayHello(ctx, &pb.NoParam{}): Calls the SayHello RPC method on the GreetServiceClient. The request sent is of type NoParam, meaning the request carries no parameters.
4. Error Handling (if err != nil): Checks if the gRPC call returned an error, and logs the error message if it failed.
5. log.Printf("%s", res.Message): Logs the message received in the HelloResponse from the server. The %s formats the message as a string.
*/
