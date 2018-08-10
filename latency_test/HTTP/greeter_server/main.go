package main

import ( "fmt";"log" ;"net/http"; "time")



func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start:=time.Now()
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Hello World",)
		t:= time.Now()
		execution_time:= t.Sub(start)
		fmt.Fprintf(w, "execution time:%s", execution_time)
	})
	log.Fatal(http.ListenAndServe(":8090", nil))

}
