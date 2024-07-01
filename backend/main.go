package main 

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weather", weatherHandler)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}