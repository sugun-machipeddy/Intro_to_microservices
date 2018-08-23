package main

import ( "log"
          //"encoding/binary"
          "net/http"
          "io/ioutil"
          "crypto/tls"
          "golang.org/x/net/http2"
          "time"
          "strings"
          "os"
          "fmt"
          "image"
          "bytes"
          "image/png"
          "strconv"
          "encoding/binary"
         )

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
  for j := 1; j<5; j++ {

  name := strconv.Itoa(j)
  var payload int

  log.Printf("itr:%d\n",j)
  //var itr_latency [50]time.Duration
  var Avg_latency time.Duration


  url := "https://localhost:8080/?itr=" + name


  for i := 0; i < 50; i++ {
  start := time.Now()
  resp, err := hc.Get(url)
	if err != nil {
		log.Fatal(err)
	 }
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
  //log.Printf("%d", contents)
  //log.Printf("%d\n",contents)
  //log.Printf("%d", len(contents))
  //l := len(contents)
  t := time.Now()

	//log.Printf(string(contents))
  payload = binary.Size(contents)

  a := strings.SplitAfter(string(contents), ":")
  l := len(a)
  //log.Printf("%s",a[l-1])
  execution_time, _ := time.ParseDuration(a[l-1])

  response_time := t.Sub(start)


  img, _, _ := image.Decode(bytes.NewReader(contents))

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
