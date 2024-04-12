package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello another World"))
}

func main() {
	http.HandleFunc("/", Handler)

	log.Println("Start HTTP server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
