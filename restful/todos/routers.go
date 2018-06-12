package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jkak/learn-go/restful/logger"
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
		var handler http.Handler
		handler = r.HandlerFunc
		handler = logger.Logger(handler, r.Name)

		router.Methods(r.Method).Path(r.Pattern).Name(r.Name).Handler(handler)
	}
	return router
}

var routes = Routes{
	Route{"HandleRoot", "GET", "/", handleRoot},
	Route{"TodoIndex", "GET", "/todos", todoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", todoShow},
}
