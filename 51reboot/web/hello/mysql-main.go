package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var (
	// cookie by sessions
	// store = sessions.NewCookieStore([]byte("sth-very-secret"))
	store = sessions.NewFilesystemStore("sessions")
)

var n = 0

func init() {
	n += 10
}

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
		http.Redirect(w, r, "/hello", 302)
	} else {
		fmt.Fprintf(w, "user:%s, password:%s login error", user, passwd)
	}
}

// Hello for hello
func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello http\n")
}

// Login usage
func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", nil)
}

func main() {
	log.Print(n)
	db, err := sql.Open("mysql", "golang:golang@tcp(59.110.12.72:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	row := db.QueryRow("SELECt CURRENT_USER()")
	if err != nil {
		log.Fatal(err)
	}
	var user string
	row.Scan(&user)
	log.Print(user)

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		id      int
		name    string
		passwd  string
		note    string
		isadmin int
	)
	for rows.Next() {
		rows.Scan(&id, &name, &passwd, &note, &isadmin)
		log.Printf("%d %s %s %s %d", id, name, passwd, note, isadmin)
	}

	return

	// 声明式挂载
	http.HandleFunc("/login", Login)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/checkLogin", CheckLogin)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
