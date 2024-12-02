package main

import (
	"golang-http-crud/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
