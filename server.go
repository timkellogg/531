package main

import (
	"log"
	"net/http"
)

func main() {
	// initialize db
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
