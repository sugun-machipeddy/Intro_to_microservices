package main

import (
	"log"
	"os"
	"time"
	"image"
	"bytes"
	"image/png"
	"fmt"
	"google.golang.org/grpc/credentials"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "./proto"
	"strconv"
	"strings"
	"encoding/binary"
)

const (
	address     = "localhost:3000"
	//defaultName = "Hello"
)

func main() {


	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
    log.Fatalf("cert load error:: %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTransferClient(conn)

	file1, _ := os.Create("payload.txt")
  file2, _ := os.Create("latency.txt")

	defer file1.Close()
  defer file2.Close()
  for j := 1; j<5; j++ {

	name := strconv.Itoa(j)
  var payload int

	log.Printf("itr:%d\n",j)
	//var itr_latency [50]time.Duration
  var Avg_latency time.Duration

	for i := 0; i < 50; i++ {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()
	r, err := c.Upload(ctx, &pb.HelloRequest{Message: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	bytes2, _ := r.Recv()
	t := time.Now()
	//log.Printf("%d\n", bytes2.Content)
	//log.Printf("%d", len(bytes2.Content))
	payload = binary.Size(bytes2.Content)

	response_time := t.Sub(start)

	img, _, _ := image.Decode(bytes.NewReader(bytes2.Content))
	execution_time, _ := time.ParseDuration(bytes2.ExecutionTime)

	x := "received"+name+".png"
  out, err := os.Create(x)

   	if err != nil {
             fmt.Println(err)
             os.Exit(1)
   	}

  err = png.Encode(out, img)

   	if err != nil {
            fmt.Println(err)
            os.Exit(1)
   	}

		//log.Printf("execution time: %s", execution_time)
	  //log.Printf("Response time: %s", response_time)

	  latency := response_time - execution_time
	  //log.Printf("latency: %s", latency)
		Avg_latency = Avg_latency + latency
	  //itr_latency[i] = latency
	}

	log.Printf("payload size:%d bytes", payload)
	Avg_latency = Avg_latency/50
	log.Printf("Avg latency %s", Avg_latency)
	fmt.Fprintf(file1, "%d \n", payload)
	AvgLatency := strings.Trim(Avg_latency.String(), "ms")
	fmt.Fprintf(file2, "%s \n", AvgLatency)
}
}
