package main

import (
	"log"
	"net"
	"time"
	//"image"
	"os"
	//"bytes"
	"bufio"
	"fmt"
	"google.golang.org/grpc/credentials"
	//context "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "./proto"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3000"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Upload(in *pb.HelloRequest, st pb.Transfer_UploadServer) (err error) {
	start := time.Now()
	time.Sleep(50 * time.Millisecond)
	a := in.Message
	image1 := "img"+a+".png"
	file, err := os.Open(image1)

  	if err != nil {
          fmt.Println(err)
          os.Exit(1)
  	}

	defer file.Close()

	fileInfo, _ := file.Stat()
  var size int64 = fileInfo.Size()
	bytes1 := make([]byte, size)
	buffer := bufio.NewReader(file)
  _, err = buffer.Read(bytes1)

	resp := &pb.Chunk{}
	resp.Content = bytes1
	t := time.Now()
	execution_time := t.Sub(start).String()
	resp.ExecutionTime = execution_time
	if err := st.Send(resp); err != nil {
			return err
		}

return nil
	//return &pb.HelloReply{Message: execution_time}, nil
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
	pb.RegisterTransferServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
