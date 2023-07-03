package main

import (
	"log"
	"net/http"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", chatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
