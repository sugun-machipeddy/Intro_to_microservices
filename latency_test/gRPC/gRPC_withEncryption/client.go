package main

import (
	"log"
	"os"
	"time"
	"google.golang.org/grpc/credentials"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:3000"
	defaultName = "world"
)

func main() {


	// Set up a connection to the server.
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
if err != nil {
    log.Fatalf("cert load error:: %v", err)
}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
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

	var itr_latency [50]time.Duration
  var Avg_latency time.Duration

  for i := 0; i < 50; i++ {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	start := time.Now()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	t := time.Now()
	log.Printf("%d", i)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("execution time: %s", r.Message)
	//log.Printf("%T", r.Message)
	execution_time, _ := time.ParseDuration(r.Message)

  response_time := t.Sub(start)
  log.Printf("response time: %s", response_time)

	latency := response_time - execution_time
	log.Printf("latency: %s", latency)
	Avg_latency = Avg_latency + latency
  itr_latency[i] = latency
	}

	Avg_latency = Avg_latency/50
  log.Printf("Avg latency %s", Avg_latency)
}
