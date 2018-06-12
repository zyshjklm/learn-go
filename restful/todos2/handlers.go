package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	//todos := Todos{
	//	Todo{Name: "Write presetation"},
	//	Todo{Name: "Host meetup"},
	//}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintf(w, "Todo Show:%s", todoID)
}

func todoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*1024))
	if err != nil {
		log.Print(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Print(err)
		return
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err1 := json.NewEncoder(w).Encode(err); err1 != nil {
			log.Print(err1)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Print(err)
		return
	}

}
