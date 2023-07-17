package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Adds to DefaultServeMux
	// Note: you can use anonymous functions
	http.HandleFunc("/hello", handleHello)
	
	fmt.Println("Listening at :8080...")
	
	// Address and handler; handler = nil means DefaultServeMux
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Will never reach
	fmt.Println("Exiting.")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from helloapi!\n"))
}
