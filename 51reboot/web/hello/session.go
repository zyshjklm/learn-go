package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("sth-secret-key"))
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

	if name != "admin" || passwd != "admin" {
		fmt.Fprintf(w, "name:%s,password:%s, login error\n", name, passwd)
	}
	store.MaxAge(10)
	session, err := store.Get(r, "web")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(session)

	session.Values["user"] = name
	session.Save(r, w)
	http.Redirect(w, r, "/hello", 302)
}

// Hello usage:
// curl localhost:8090/hello
func Hello(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "web")
	user, ok := session.Values["user"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	}
	fmt.Fprintf(w, "hello %s", user)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/checkLogin", CheckLogin)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
