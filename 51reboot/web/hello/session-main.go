package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/sessions"
)

var (
	// cookie by sessions
	// store = sessions.NewCookieStore([]byte("sth-very-secret"))
	store = sessions.NewFilesystemStore("sessions")
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
		sesn, err := store.Get(r, "web")
		if err != nil {
			log.Print(err)
			return
		}
		// sesn := sessions.NewSession(store, "web")
		sesn.Values["user"] = user
		sesn.Save(r, w)
		http.Redirect(w, r, "/hello", 302)
	} else {
		fmt.Fprintf(w, "user:%s, password:%s login error", user, passwd)
	}
}

// Hello for hello
func Hello(w http.ResponseWriter, r *http.Request) {
	sesn, _ := store.Get(r, "web")
	_, ok := sesn.Values["user"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user := sesn.Values["user"]
	hello := "Hello " + user.(string) + "\n"
	io.WriteString(w, hello)
}

// result: Hello user=admin
// 在开发者工具的Network中，点击hello项，在右侧的Cookies中能看到一对user:admin的key:value

// Login usage
func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", nil)
}

func main() {
	// 声明式挂载
	http.HandleFunc("/login", Login)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/checkLogin", CheckLogin)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
