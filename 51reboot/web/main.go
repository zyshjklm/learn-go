package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func render(w http.ResponseWriter, name string, data interface{}) {
	tplFile := filepath.Join("template", name+".tpl")
	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Login usage
func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", nil)
}

func main() {
	http.HandleFunc("/login", Login)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
