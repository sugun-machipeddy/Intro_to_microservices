package main

import (
	"log"
	"time"
	"fmt"
	"os"
	"encoding/binary"
	"strconv"
	"strings"
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

	file1, _ := os.Create("payload.txt")
  file2, _ := os.Create("latency.txt")

	defer file1.Close()
  defer file2.Close()

  for i := 1; i < 51; i++ {

	name := strconv.Itoa(i)
	var itr_latency [51]time.Duration
  var Avg_latency time.Duration
	var payload int

	log.Printf(" count:%d", i)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for j := 0; j< 50; j++{

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	start := time.Now()
	r, _ := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	t := time.Now()


	a := strings.SplitAfter(string(r.Message), ":")

  execution_time, _ := time.ParseDuration(a[1])

	//log.Printf("execution time: %s", execution_time)
	//log.Printf("%T", r.Message)

	var b = []byte(r.Message)

	//fmt.Printf("length of payload:%d \n", len(b))
	//log.Printf("payload: %s", string(b))
	//log.Printf("memory size of payload:%d\n", binary.Size(b))
	payload = binary.Size(b)
	//fmt.Printf("Capacity:%d ", b.Cap())


  response_time := t.Sub(start)
  //log.Printf("response time: %s", response_time)

	latency := response_time - execution_time
	//log.Printf("latency: %s", latency)
	Avg_latency = Avg_latency + latency
  itr_latency[i] = latency

	}
	log.Printf("payload size:%d bytes", payload)
	Avg_latency = Avg_latency/50
  log.Printf("Avg latency %s", Avg_latency)
  fmt.Fprintf(file1, "%d \n", payload)
  AvgLatency := strings.Trim(Avg_latency.String(), "ms")
  fmt.Fprintf(file2, "%s \n", AvgLatency)
}
}
