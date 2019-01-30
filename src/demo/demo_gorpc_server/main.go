// demo_gorpc_server project main.go

package main

import (
	"log"
	"net"

	pb "pb/helloworld"

	//	proto "github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{C: in.A + in.B}, nil
}

func (s *server) Sub(ctx context.Context, in *pb.SubRequest) (*pb.SubReply, error) {
	return &pb.SubReply{C: in.A - in.B}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
