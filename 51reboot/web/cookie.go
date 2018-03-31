package main

import (
	"fmt"
	"html/template"
	"io"
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

// Login usage
func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", nil)
}

// CheckLogin usage
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("user")
	passwd := r.FormValue("password")

	if name == "admin" && passwd == "admin" {
		cookie := &http.Cookie{
			Name:   "user",
			Value:  name,
			MaxAge: 10,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/hello", 302)
	} else {
		fmt.Fprintf(w, "name:%s,password:%s, login error\n", name, passwd)
	}
}

// Hello usage:
// curl localhost:8090/hello
func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello http\n")
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/checkLogin", CheckLogin)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
