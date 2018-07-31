package main

import ( "fmt";"log" ;"net/http"; "io/ioutil")

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Secondpage!")
	resp, err := http.Get("http://api3-service.default.svc.cluster.local")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()	
	contents, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(contents))
 	fmt.Println("Endpoint Hit: Secondpage")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
	handleRequests()
}

