package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Todo Index!")
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintf(w, "Todo Show:%s", todoID)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoot)
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoId}", todoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
curl localhost:8080/todos
Todo Index!%

curl localhost:8080/todos/23
Todo Show:23%

curl localhost:8080/todos/ab
Todo Show:ab%

*/
