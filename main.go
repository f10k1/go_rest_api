package main

import (
	"log"
	"net/http"
	"rest_api/handlers"
	"rest_api/middlewares"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var listOfMiddlewares = []Middleware{
	middlewares.TokenAuthMiddleware,
}

func main() {
	var handler http.HandlerFunc = handlers.HandleClientProfile
	for _, middleware := range listOfMiddlewares {
		handler = middleware(handler)
	}

	http.HandleFunc("/user/profile", handler)

	log.Println("Server working on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
