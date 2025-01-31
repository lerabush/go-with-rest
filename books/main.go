package main

import (
	"books/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)

	fmt.Println("Server listening on port 8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
