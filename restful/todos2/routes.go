package main

import (
	"net/http"
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

var routes = Routes{
	Route{"HandleRoot", "GET", "/", handleRoot},
	Route{"TodoIndex", "GET", "/todos", todoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", todoShow},
	Route{"TodoCreate", "POST", "/todos", todoCreate},
}
