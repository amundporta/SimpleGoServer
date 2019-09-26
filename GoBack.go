package main

import (
	"fmt"
	"log"
	"net/http"
)

// server html

// ta imot verdi i JSON format
// legg i liste

func main() {
	http.Handle("/", http.HandlerFunc(handleFunc))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test en to tre %s", r.URL.Path[1:])
}
