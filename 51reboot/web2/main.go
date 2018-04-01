package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	age    = 300
	dbName = "web"
)

var (
	db    *sqlx.DB
	store = sessions.NewCookieStore([]byte("sth-secret-key"))
	fname = flag.String("f", "database/user.db", "file name of sqlite3")
)

// NeedLogin for check login
func NeedLogin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "web")
		_, ok := session.Values["user"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		h(w, r)
	}
}

func main() {
	flag.Parse()
	store.MaxAge(age)

	var err error
	db, err = sqlx.Open("sqlite3", *fname)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// loginCounter := NewCounter(http.HandlerFunc(Login))
	// http.Handle("/login", loginCounter)
	// http.Handle("/", loginCounter)
	// http.HandleFunc("/loginCounter", loginCounter.GetCounter)
	http.Handle("/login", http.HandlerFunc(Login))
	http.Handle("/", http.HandlerFunc(Login))

	http.HandleFunc("/checkLogin", CheckLogin)

	// /add render() html; post action to /create
	http.HandleFunc("/add", NeedLogin(RenderAdd))
	http.HandleFunc("/create", NeedLogin(Create))

	// /delete
	http.HandleFunc("/delete", NeedLogin(Delete))

	// list
	http.HandleFunc("/list", NeedLogin(List))

	// modify
	http.HandleFunc("/update", NeedLogin(RenderUpdate))
	http.HandleFunc("/modify", NeedLogin(Modify))

	// file server
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	// api
	http.HandleFunc("/users", Users)

	// middleware
	mux := handlers.LoggingHandler(os.Stderr, http.DefaultServeMux)
	c := NewCounter(mux)
	http.HandleFunc("/counter", c.GetCounter)

	log.Fatal(http.ListenAndServe(":8090", c))
}
