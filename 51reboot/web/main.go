package main

import (
	"fmt"
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
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// CheckLogin usage:
// curl http://localhost:8090/checkLogin?user=admin&password=admin
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("user")
	passwd := r.FormValue("password")
	if user == "admin" && passwd == "admin" {
		fmt.Fprintf(w, "login ok")
	} else {
		fmt.Fprintf(w, "user:%s, password:%s", user, passwd)
	}
}

// Login usage
func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", nil)
}

func main() {
	// 声明式挂载
	http.HandleFunc("/login", Login)
	http.HandleFunc("/checkLogin", CheckLogin)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
