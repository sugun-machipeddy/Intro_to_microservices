package main

import ( "log"; "net/http"; "io/ioutil"; "crypto/tls"; "golang.org/x/net/http2"; "time" )

func main(){

  hc := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

  start := time.Now()
  resp, err := hc.Get("https://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	log.Printf(string(contents))
  t := time.Now()
  elapsed := t.Sub(start)
  log.Printf("Duration: %s", elapsed)
}
