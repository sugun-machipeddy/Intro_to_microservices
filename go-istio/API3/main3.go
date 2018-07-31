package main

import ( "fmt";"log" ;"net/http")

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Thirdpage!")
 	fmt.Println("Endpoint Hit: Thirdpage")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
	handleRequests()
}

