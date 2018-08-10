/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"log"
	"os"
	"time"

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
