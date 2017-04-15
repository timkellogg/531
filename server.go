package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// initialize db
	router := NewRouter()
	fmt.Println("Serve up on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
