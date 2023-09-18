package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func handle(w http.ResponseWriter, r *http.Request) {
	spew.Dump(r.Header) //port := 9999
	//fmt.Printf("Server is listening on port %d...\n", port)
	//http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	fmt.Fprintf(w, "Hello 👋")

}

func main() {
	http.HandleFunc("/hello", handle)
	log.Fatal(http.ListenAndServe(":8090", nil))

}
