package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route for router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes for route slice
type Routes []Route

// NewRouter to new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		router.Methods(r.Method).Path(r.Pattern).Name(r.Name).Handler(r.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{"HandleRoot", "GET", "/", handleRoot},
	Route{"TodoIndex", "GET", "/todos", todoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", todoShow},
}
