package main

import ( "log"; "net/http"; "io/ioutil"; "crypto/tls"; "golang.org/x/net/http2"; "time"; "strings" )

func main(){

  hc := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
  var itr_latency [50]time.Duration
  var Avg_latency time.Duration

  for i := 0; i < 50; i++ {
  start := time.Now()
  resp, err := hc.Get("https://localhost:8080")
	if err != nil {
		log.Fatal(err)
	 }
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
  t := time.Now()
  log.Printf("%d", i)
	log.Printf(string(contents))

  a := strings.SplitAfter(string(contents), ":")
  execution_time, _ := time.ParseDuration(a[1])
  log.Printf("execution time: %s", execution_time)

  response_time := t.Sub(start)
  log.Printf("Response time: %s", response_time)

  latency := response_time - execution_time
  log.Printf("latency: %s", latency)
  Avg_latency = Avg_latency + latency
  itr_latency[i] = latency
  }


  Avg_latency = Avg_latency/50
  log.Printf("Avg latency %s", Avg_latency)
}
