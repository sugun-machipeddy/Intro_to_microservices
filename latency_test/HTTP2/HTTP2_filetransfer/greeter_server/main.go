package main

import ( "fmt"
					"log"
				 	//"math/rand"
				 	//"strconv"
					"net/http"
				 	"golang.org/x/net/http2"
					"time"
					"os"
					"bufio"
					"encoding/binary"
				)


func main(){
	var s http.Server
	http2.VerboseLogs = true
	s.Addr = ":8080"

	http2.ConfigureServer(&s, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    start:=time.Now()
		q := r.URL.Query()
		a := q.Get("itr")

		w.Header().Set("Content-Type", "image/png")
		//time.Sleep(50 * time.Millisecond)

		image1 := "img"+a+".png"
		file, err := os.Open(image1)
		file1, _ := os.Create("Actualpayload.txt")
	  	if err != nil {
	          fmt.Println(err)
	          os.Exit(1)
	  	}

		defer file.Close()

		fileInfo, _ := file.Stat()
	  var size int64 = fileInfo.Size()
		bytes1 := make([]byte, size)
		buffer := bufio.NewReader(file)
	  _, err = buffer.Read(bytes1)
		t:= time.Now()

		payload := binary.Size(bytes1)
		fmt.Fprintf(file1, "%d \n", payload)
		execution_time:= t.Sub(start).String()
		fmt.Fprintf(w, "%s:%s", bytes1, execution_time)
	})

	log.Fatal(s.ListenAndServeTLS("server_http2.crt", "server.key"))

}
