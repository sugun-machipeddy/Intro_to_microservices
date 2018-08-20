package main

import ( "log"; "strconv"; "encoding/binary"; "net/http"; "io/ioutil"; "crypto/tls"; "golang.org/x/net/http2"; "time"; "strings"; "os"; "fmt" )

func main(){

  hc := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}

  file1, _ := os.Create("payload.txt")
  file2, _ := os.Create("latency.txt")

	defer file1.Close()
  defer file2.Close()

  for i := 1; i < 51; i++ {

  name := strconv.Itoa(i)
  var itr_latency [51]time.Duration
  var Avg_latency time.Duration
  var payload int
  url := "https://localhost:8080/?itr=" + name

  log.Printf("%d", i)

  for j := 0; j< 50; j++{

  start := time.Now()
  resp, err := hc.Get(url)
	if err != nil {
		log.Fatal(err)
	 }
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
  t := time.Now()

	//log.Printf(string(contents))

  a := strings.SplitAfter(string(contents), ":")
  execution_time, _ := time.ParseDuration(a[1])

  var b = []byte(string(contents))
  payload = binary.Size(b)
  //log.Printf("payload size:%d bytes", payload)

  //log.Printf("execution time: %s", execution_time)

  response_time := t.Sub(start)
  //log.Printf("Response time: %s", response_time)

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
