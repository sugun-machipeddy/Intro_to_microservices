package main

import ( "fmt"; "log" ;"net/http"; "golang.org/x/net/http2"; "time")


func main(){
	var s http.Server
	http2.VerboseLogs = true
	s.Addr = ":8080"

	http2.ConfigureServer(&s, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    start:=time.Now()
		w.Header().Set("Content-Type", "text/plain")
		time.Sleep(50 * time.Millisecond)
		t:= time.Now()
		execution_time:= t.Sub(start)
		fmt.Fprintf(w, "execution time:%s", execution_time)
	})

	log.Fatal(s.ListenAndServeTLS("server.crt", "server.key"))

}
