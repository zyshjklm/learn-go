package main 

import (
	"fmt"
	"log"
	"net/http"
	// "strings"
)

// type Hello struct{}


// func (h Hello) ServeHTTP(
// 	w http.ResponseWriter,
// 	r *http.Request) {
// 	fmt.Fprint(w, "Hello!")
// }

////////////////////////////////////

type String string

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, str + "\n")
}

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (st Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	tmp := st.Greeting+" "+st.Punct+" "+st.Who + "\n"
	fmt.Fprint(w, tmp)
}


func main() {
	str := String("I'm a frayed knot.")
	stc := &Struct{"Hello", ":", "Gophers!"}

	http.Handle("/string", str)
	http.Handle("/struct", stc)

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
