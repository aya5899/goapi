package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// define handler
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// return "string" whatever received
		io.WriteString(w, "Hello, world!\n")
	}
	// register helloHandler
	http.HandleFunc("/hello", helloHandler)

	log.Println("server start at port 8080")
	// launch server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
