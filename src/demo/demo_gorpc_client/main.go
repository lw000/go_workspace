// demo_gorpc_client project main.go
package main

import (
	"log"
	"os"

	//	pb "demo_gorpc_client/pb/helloworld"
	pb "pb/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	{
		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting SayHello: %s", r.Message)
	}

	{
		r, err := c.Add(context.Background(), &pb.AddRequest{A: 1, B: 2})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting Add: %s", r.C)
	}

	{
		r, err := c.Sub(context.Background(), &pb.SubRequest{A: 1, B: 2})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting Sub: %s", r.C)
	}
}
