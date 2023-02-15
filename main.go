package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("listening at localhost on port  :9090")

	log.Fatal(http.ListenAndServe(":9090", nil))
}
