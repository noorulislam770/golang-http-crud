package router

import (
	"golang-http-crud/handlers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// User CRUD APIs
	mux.HandleFunc("/users", handlers.UserHandler)
	mux.HandleFunc("/users/search", handlers.SearchHandler)

	return mux
}
