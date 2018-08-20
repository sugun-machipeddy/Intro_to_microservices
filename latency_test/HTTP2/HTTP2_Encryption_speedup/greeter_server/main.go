package main

import ( "fmt"; "log"; "math/rand"; "strconv" ;"net/http"; "golang.org/x/net/http2"; "time")


func main(){
	var s http.Server
	http2.VerboseLogs = true
	s.Addr = ":8080"

	http2.ConfigureServer(&s, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    start:=time.Now()
		w.Header().Set("Content-Type", "text/plain")
		time.Sleep(50 * time.Millisecond)
		q := r.URL.Query()
		a := q.Get("itr")
		n, _ := strconv.Atoi(a)

		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, (n*40000))
	  for i := range b {
	      b[i] = letterBytes[rand.Intn(len(letterBytes))]
	  }

		t:= time.Now()
		execution_time:= t.Sub(start)
		fmt.Fprintf(w, "%s:%s", string(b), execution_time)
	})

	log.Fatal(s.ListenAndServeTLS("server_http2.crt", "server.key"))

}
