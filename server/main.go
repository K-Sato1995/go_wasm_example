package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Run Web Server on http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../assets"))))
}
