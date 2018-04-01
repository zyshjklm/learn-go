package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	_  http.Handler = &counter{}
	// 进行编译检查，证明counter实现了Handler接口。
)

// User for sql row
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-" xml:"-"`
	Note     string `json:"note"`
	Isadmin  bool   `json:"isadmin"`
}
type counter struct {
	h     http.Handler
	count map[string]int
}

// NewCounter new a counter with Handler
func NewCounter(h http.Handler) *counter {
	return &counter{
		h:     h,
		count: make(map[string]int),
	}
}

func (c *counter) GetCounter(w http.ResponseWriter, r *http.Request) {
	for path, count := range c.count {
		fmt.Fprintf(w, "%s\t%d\n", path, count)
	}
}
func (c *counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.count[r.URL.Path]++
	c.h.ServeHTTP(w, r)
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
	render(w, "list.html", users)
}

// NeedLogin for check login
func NeedLogin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("user")
		if err != nil {
			render(w, "login", "login out of time")
			return
		}
		h(w, r)
	}
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
		MaxAge: 10,
	})
	http.Redirect(w, r, "/list", 302)
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
	// var err error
	// db, err = sqlx.Open("mysql", "golang:golang@tcp(59.110.12.72:3306)/go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// http.HandlerFunc -> func,  -> use http.HandleFunc to mount
	// http.Handler -> interface, -> use http.Handle to mount
	// Login -> http.Handler :
	// http.HandlerFunc(Login) -> http.Handler

	// 声明式挂载
	loginCounter := NewCounter(http.HandlerFunc(Login))
	http.Handle("/login", loginCounter)
	http.HandleFunc("/loginCounter", loginCounter.GetCounter)
	// 专属counter
	http.HandleFunc("/add", NeedLogin(Add))
	http.HandleFunc("/list", NeedLogin(List))
	http.HandleFunc("/hello", NeedLogin(Hello))
	http.HandleFunc("/checkLogin", CheckLogin)

	h := handlers.LoggingHandler(os.Stderr, http.DefaultServeMux)
	c := NewCounter(h)
	http.HandleFunc("/counter", c.GetCounter)
	// 全局counter

	log.Fatal(http.ListenAndServe(":8090", c))
	// 对于一个/login请求，会被多次中间件封装：
	// /login -> c(counter) -> h(log) ->mux route --> // couter + login //
}
