package main

import (
	"log"
	"net"
	"time"
	"google.golang.org/grpc/credentials"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3000"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	start := time.Now()
	time.Sleep(50 * time.Millisecond)
	t := time.Now()
	execution_time := t.Sub(start).String()

	return &pb.HelloReply{Message: execution_time}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
if err != nil {
    log.Fatalf("Failed to setup tls: %v", err)
}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
