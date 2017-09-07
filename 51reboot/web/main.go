package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

// User for sql row
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-" xml:"-"`
	Note     string `json:"note"`
	Isadmin  bool   `json:"isadmin"`
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

// Add to add a user
func Add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	passwd := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))
	note := r.FormValue("note")

	res, err := db.Exec("INSERT INTO user VALUES(NULL, ?, ?, ?, ?)", name, passwd, note, 1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Print(res.LastInsertId())
	log.Print(res.RowsAffected())
}

// List all user as table
func List(w http.ResponseWriter, r *http.Request) {
	var users []User
	err := db.Select(&users, "SELECT * FROM user")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	render(w, "list", users)
}

// CheckLogin usage
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("user")
	passwd := r.FormValue("password")
	// 查询数据库
	var user User
	err := db.Get(&user, "SELECT password FROM user WHERE name = ?", name)
	if err != nil {
		render(w, "login", "user not found")
		return
	}
	// 计算passwd的md5与刚才获取的值是否相同。不同则失败。
	if fmt.Sprintf("%x", md5.Sum([]byte(passwd))) != user.Password {
		render(w, "login", "bad password or username")
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "user",
		Value:  name,
		MaxAge: 30,
	})
	http.Redirect(w, r, "/hello", 302)
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
	var err error
	db, err = sqlx.Open("mysql", "golang:golang@tcp(59.110.12.72:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 声明式挂载
	http.HandleFunc("/login", Login)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/list", List)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/checkLogin", CheckLogin)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
